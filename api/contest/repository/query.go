package repository

const queryRegisterCompetition = `
	INSERT INTO
		participants
	(
		contest_id,
		name,
		birth,
		category,
		institution,
		email,
		instagram,
		line,
		phone_number,
		video_url,
		payment_proof,
		form
	) VALUES (
		:contest_id,
		:name,
		:birth,
		:category,
		:institution,
		:email,
		:instagram,
		:line,
		:phone_number,
		:video_url,
		:payment_proof,
		:form
	)
`

const queryUpdateParticipant = `
	UPDATE
		participants
	SET
		%s
	WHERE
		participants.id = :id
`

const queryGetParticipants = `
	SELECT
		participants.id,
		participants.name,
		participants.birth,
		participants.category,
		participants.status,
		participants.institution,
		participants.email,
		participants.instagram,
		participants.line,
		participants.phone_number,
		participants.form,
		participants.video_url,
		participants.payment_proof,
		contests.name AS contest_name
	FROM
		participants
	JOIN
		contests ON participants.contest_id = contests.id
	%s
`
