package repository

const queryRegisterCompetition = `
	INSERT INTO
		participants
	(
		contests_id,
		name,
		origin,
		phone_number,
		video_url,
		payment_proof,
		form_url
	) VALUES (
		:contest_id,
		:name,
		:origin,
		:phone_number,
		:video_url,
		:payment_proof,
		:form_url
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
		participants.status,
		participants.origin,
		participants.phone_number,
		participants.form_url,
		participants.video_url,
		participants.payment_proof,
		contests.name AS contest_name
	FROM
		participants
	JOIN
		contests ON participants.contests_id = contests.id
	%s
`
