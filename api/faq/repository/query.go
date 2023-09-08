package repository

const queryAddFaq = `
	INSERT INTO faq (
		category,
		title,
		question,
		answer
	) VALUES (
		:category,
		:title,
		:question,
		:answer
	)
`

const queryGetFaq = `
	SELECT
		id,
		category
		title,
		question,
		answer
	FROM faqs
		%s
`

const queryDeleteFaq = `
	DELETE FROM
		faqs
	WHERE
		id = :id
`

const queryUpdateFaq = `
	UPDATE faq
	SET
		category = :category,
		title = :title,
		question = :question,
		answer = : answer
	WHERE
		id = :id
`
