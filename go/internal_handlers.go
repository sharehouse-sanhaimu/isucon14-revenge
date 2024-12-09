package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// MEMO: 一旦最も待たせているリクエストに適当な空いている椅子マッチさせる実装とする。おそらくもっといい方法があるはず…
	ride := &Ride{}
	// 最も待たせているride（まだchair_idがnull）の、最も古い(ORDER BY created_at)ものを取得
	if err := db.GetContext(ctx, ride, `SELECT * FROM rides WHERE chair_id IS NULL ORDER BY created_at LIMIT 1`); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	matched := &Chair{}
	for i := 0; i < 10; i++ {
		// Complatedのchairsの中で最も古いものを取得
		query := `SELECT * FROM chairs LEFT JOIN rides ON chairs.id = rides.chair_id LEFT JOIN ride_statuses ON ride_statuses.ride_id = rides.id WHERE chairs.is_active = TRUE GROUP BY chairs.id HAVING COUNT(CASE WHEN ride_statuses.status != 'COMPLETED' THEN 1 END) = 0 ORDER BY chairs.updated_at LIMIT 1`
		err := db.GetContext(ctx, matched, query)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// 結果がない場合（ErrNoRows）、ループを継続
				continue
			}
			// それ以外のエラーは500エラーとして返す
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		// クエリ結果が得られた場合、ループを抜ける
		break
	}

	// ループが終了した後、結果が得られたかを確認
	if matched.ID == "" {
		// クエリ結果が得られなかった場合の処理
		writeError(w, http.StatusNotFound, fmt.Errorf("no completed chairs found"))
		return
	}

	// マッチングしたrideとchairを紐付ける
	if _, err := db.ExecContext(ctx, "UPDATE rides SET chair_id = ? WHERE id = ?", matched.ID, ride.ID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
