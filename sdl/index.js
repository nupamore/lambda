const mysql = require('mysql2/promise')

const query = `
    SELECT target
    FROM simple_dynamic_link
    WHERE link_id = ?;
`

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
        const [rows] = await connection.query(query, event.params.path.linkId)
        if (!rows.length) {
            response.statusCode = 404
            response.body = JSON.stringify(`Couldn't find target`)
        } else {
            const { target } = rows[0]
            response.statusCode = 302
            response.headers = { Location: target }
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