INSERT INTO permissions (id, name, is_admin) VALUES
	(UUID(), "admin", true),
 	(UUID(), "vedouci", false),
	(UUID(), "zamestnanec", false),
	(UUID(), "brigadnik", false);

INSERT INTO users (id, user_name, password, first_name, last_name, permission_id) VALUES
	(UUID(), "admin", "admin", NULL, NULL, (SELECT id From permissions WHERE name = "admin")),
	(UUID(), "sadlof", "heslo", "Franta", "Sadlo", (SELECT id From permissions WHERE name = "vedouci")),
	(UUID(), "maslol", "heslo", "Lojza", "Maslo", (SELECT id From permissions WHERE name = "zamestnanec")),
	(UUID(), "pazitkap", "heslo", "Pepa", "Pazitka", (SELECT id From permissions WHERE name = "brigadnik"));

INSERT INTO firms (id, name, email, tel_number, description) VALUES
	(UUID(), "SoftCorp s.r.o.", "soft@corp.cz", "444555666", "Nase firma. Pro pridavani internich projektu."),
	(UUID(), "Google a.s.", "google@gmail.com", "123456789", "Proste Google. Co vice k tomu rict. Spolehlivy zakaznici kteri vcas plati."),
	(UUID(), "Seznam s.r.o", "seznam@seznam.cz", "987654321", NULL),
	(UUID(), "Alza.cz", "alza@alza.cz", NULL, NULL),
	(UUID(), "ABRA a,s,", "info@abra.cz", NULL, NULL),
	(UUID(), "Moje Firma s.r.o", "moje@firma.cz", NULL, NULL),
	(UUID(), "Cherry GMbh", "cherry@cehrry.cz", NULL, NULL),
	(UUID(), "Uz me nic nenapada", "mail@seznam.cz", NULL, NULL),
	(UUID(), "Dalsi firma", "firmas@seznam.cz", NULL, NULL),
	(UUID(), "And another one", "ten@email.com", NULL, NULL),
	(UUID(), "Tata a syn", "lol@wut.kappa", NULL, NULL);

INSERT INTO projects (id, name, code, description, start_date, user_id, firm_id) VALUES
	(UUID(), "Naplnit ISSZP pocatecnimy daty", "ISSZP-Init", "ISSZP musi byt naplneno daty pred prvotnim uvedenim do provozu", NOW(),
		(SELECT id FROM users WHERE user_name = "sadlof"),
		(SELECT id FROM firms WHERE name = "SoftCorp s.r.o.")),
	(UUID(), "Testovai ISSZP", "ISSZP-Test", "Je nutne aby Lojza Maslo s Pepou Pazitkou otestovali poradne ISSZP aplikaci", NOW(),
		(SELECT id FROM users WHERE user_name = "maslol"),
		(SELECT id FROM firms WHERE name = "SoftCorp s.r.o."));