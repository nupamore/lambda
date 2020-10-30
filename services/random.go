package services

import "fmt"

type imageRow struct {
	ChannelID string
	FileID    string
	FileName  string
}

func imgURL(row imageRow) string {
	return fmt.Sprintf(
		"https://cdn.discordapp.com/attachments/%s/%s/%s",
		row.ChannelID, row.FileID, row.FileName,
	)
}

// GetRandomImage : get random image from guildID
func GetRandomImage(guildID string) (*string, error) {
	var image imageRow
	db, err := GetDB()
	defer db.Close()

	query := `
		SELECT channel_id, file_id, file_name
		FROM discord_images
		WHERE guild_id = ?
		ORDER BY rand() limit 1;
	`

	err = db.QueryRow(query, guildID).Scan(
		&image.ChannelID,
		&image.FileID,
		&image.FileName,
	)
	url := imgURL(image)
	return &url, err
}
