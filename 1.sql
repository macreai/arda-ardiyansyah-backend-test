SELECT b.id_murid, a.name, b.status as pendidikan_terakhir, a.time_create as time_create, b.time_create as time_update 
FROM murid a
JOIN pendidikan b ON a.id = b.id_murid;