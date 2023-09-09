package repository

const queryGetUserByEmail = `
	SELECT
		users.id
		GROUP_CONCAT(roles.name SEPARATOR ', ') AS roles
		users.provider
		users.provider_id
		users.name
		users.email
		users.created_at
		users.updated_at
	FROM
		users
	JOIN
		user_has_roles ON user_has_roles.user_id = users.id
	JOIN
		roles ON user_has_roles.role_id = roles.id
	%s
	GROUP BY id;
`

const queryUpdateUserProviderID = `
	UPDATE
		users
	SET
		users.provider_id = :provider_id
	WHERE
		users.id = :id;
`
