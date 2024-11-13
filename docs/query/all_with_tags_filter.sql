SELECT 
			tasks.id,
			tasks.text,
			tasks.due,
			IFNULL(GROUP_CONCAT(tags.tag, '; '), '') AS tags
		FROM 
			tasks
		LEFT JOIN 
			tags ON tasks.id = tags.task_id
		WHERE 
			tags.tag IN ('work', 'studdy') -- заполнить сплитом
		GROUP BY 
			tasks.id
		HAVING 
			COUNT(DISTINCT tags.tag) = 2; --Посчитать кол-во тэгов

