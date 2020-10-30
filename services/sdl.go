package services

// LinkToTarget : linkID to target url
func LinkToTarget(linkID string) (*string, error) {
	var target string
	db, err := GetDB()
	defer db.Close()

	query := `
		SELECT target
		FROM simple_dynamic_link
		WHERE link_id = ?;
	`

	err = db.QueryRow(query, linkID).Scan(&target)
	return &target, err
}
