CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE emergency_details (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL,
    bound_to_bed BOOLEAN NOT NULL DEFAULT FALSE,
    has_chair_wheel BOOLEAN NOT NULL DEFAULT FALSE,
    last_updated_method TEXT CHECK(last_updated_method IN ('manual', 'file')),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE emergency_contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    emergency_details_id INTEGER NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    relationship TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    email TEXT,
    FOREIGN KEY (emergency_details_id) REFERENCES emergency_details(id) ON DELETE CASCADE
);

CREATE VIEW users_with_emergency_details AS
SELECT  u.id,
        u.email,
        u.created_at,
        u.updated_at,
        ed.bound_to_bed,
        ed.has_chair_wheel,
        ed.updated_at,
        ed.last_updated_method,
        ec.last_name AS contact_last_name,
        ec.first_name AS contact_first_name,
        ec.relationship AS contact_relationship,
        ec.phone_number AS contact_phone_number,
        ec.email AS contact_email
FROM users u
LEFT JOIN
        emergency_details ed ON u.id = ed.user_id
LEFT JOIN 
        emergency_contacts ec ON ed.id = ec.emergency_details_id;