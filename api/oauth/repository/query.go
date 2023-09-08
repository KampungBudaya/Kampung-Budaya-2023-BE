package repository

const queryGetUserByEmail = `
	SELECT
		users.id
		users.provider
		users.provider_id
		users.name
		users.email
		users.created_at
		users.updated_at
	FROM
		users
	%s;
`

const queryUpdateUserProviderID = `
	UPDATE
		users
	SET
		users.provider_id = :provider_id
	WHERE
		users.id = :id;
`
