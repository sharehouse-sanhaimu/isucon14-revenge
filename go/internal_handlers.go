package main

import (
	"database/sql"
	"errors"
	"net/http"
)

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// MEMO: 一旦最も待たせているリクエストに適当な空いている椅子マッチさせる実装とする。おそらくもっといい方法があるはず…
	ride := &Ride{}
	if err := db.GetContext(ctx, ride, `SELECT * FROM rides WHERE chair_id IS NULL ORDER BY created_at LIMIT 1`); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	matched := &ChairAvailable{}
	if err := db.GetContext(ctx, matched, `SELECT chair_id FROM chair_available WHERE is_available = TRUE LIMIT 1`); err != nil {
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	writeError(w, http.StatusInternalServerError, err)
	return
}

	if _, err := db.ExecContext(ctx, "UPDATE rides SET chair_id = ? WHERE id = ?", matched.ChairID, ride.ID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// chair_availableテーブルのis_availableをfalseに更新
	if _, err := db.ExecContext(ctx, "UPDATE chair_available SET is_available = FALSE WHERE chair_id = ?", matched.ChairID); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
	}

	w.WriteHeader(http.StatusNoContent)
}
