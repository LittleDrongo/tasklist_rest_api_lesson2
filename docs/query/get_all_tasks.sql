SELECT 
			tasks.id,
			tasks.text,
			tasks.due,
			IFNULL(GROUP_CONCAT(tags.tag, '; '), '') AS tags
		FROM 
			tasks
		LEFT JOIN 
			tags ON tasks.id = tags.task_id
		GROUP BY 
			tasks.id