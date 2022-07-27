package dbfellow

func createTables() error {

	if _, err := Db.Exec(`create table if not exists users (
		tgID        	text primary key,
		timeSt			timestamp with time zone default now(),
		firstname   	text,
		nameInBot  		text,
		age         	text default 'не указан',
		sex         	text default 'не указан',
		photo       	bytea,
		interests   	text,
		aboutMe     	text default '',
		saved       	boolean not null default false,

		location     	text default '',

		partnersAge     text default 'не важно',
		partnersSex     text default 'не важно',

		likeIDs     	text[] default array[]::text[],
		currentR    	int default 1,
		matches     	text[] default array[]::text[],
		currentM    	int default 1
		searchResults   text[] default array[]::text[],
		currentCV       int default 1

		stage           text default '',
		mesNO           bigint default 0,	
	);`); err != nil {
		return err
	}

	if _, err := Db.Exec(`create table if not exists likes (
		likeID          serial primary key,
		liketime        timestamp with time zone default now(),
		likeFromTgID    text,
		likeToTgID      text,
		match           boolean default false
	);`); err != nil {
		return err
	}

	if _, err := Db.Exec(`create table if not exists states_ (
		owl	        	text,
		num         	bigint primary key,
		title			text default '' 
	);`); err != nil {
		return err
	}

	return nil
}
