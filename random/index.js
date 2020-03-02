const mysql = require('mysql2/promise')

const query = `
    SELECT channel_id, file_id, file_name
    FROM discord_images
    WHERE guild_id = ?
    ORDER BY rand() limit 1;
`

function imgUrl(channel_id, file_id, file_name) {
    return `https://cdn.discordapp.com/attachments/${channel_id}/${file_id}/${file_name}`
}

async function image(event) {
    const connection = await mysql.createConnection({
        host: process.env.host,
        user: process.env.user,
        password: process.env.password,
        port: process.env.port,
        database: process.env.database,
    })
    const response = {}
    try {
        const [rows] = await connection.query(query, event.params.path.guildId)
        if (!rows.length) {
            response.statusCode = 404
            response.body = JSON.stringify(`Couldn't find any image`)
        } else {
            const { channel_id, file_id, file_name } = rows[0]
            const url = imgUrl(channel_id, file_id, file_name)
            response.statusCode = 302
            response.headers = { Location: url }
        }
    } catch (err) {
        response.statusCode = 400
        response.body = JSON.stringify(err)
    }
    connection.end()
    return response
}

exports.handler = async (event) => {
    return await image(event)
}