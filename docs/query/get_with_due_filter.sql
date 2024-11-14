-- Active: 1731529367129@@127.0.0.1@3306
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
            1=1
            AND tasks.due >= '2024-11-15 00:00:00' AND tasks.due < '2024-11-16 00:00:00'
		GROUP BY 
			tasks.id