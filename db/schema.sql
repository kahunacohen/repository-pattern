-- Tables

CREATE TABLE companies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_sync_active BOOLEAN NOT NULL DEFAULT 0,
    employee_sync_ftp_address TEXT,
    employee_sync_ftp_input_dir TEXT,
    employee_sync_ftp_output_dir TEXT,
    employee_sync_ftp_password TEXT,
    employee_sync_ftp_user TEXT,
    name TEXT NOT NULL
);

CREATE TABLE end_work_period_reasons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT
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

CREATE TABLE emergency_details (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    patient_id INTEGER NOT NULL,
    bound_to_bed BOOLEAN NOT NULL DEFAULT FALSE,
    has_chair_wheel BOOLEAN NOT NULL DEFAULT FALSE,
    last_updated_method TEXT CHECK(last_updated_method IN ('manual', 'file')),
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
);

CREATE TABLE employees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    contact_info_phone_number TEXT,
    contact_info_phone_number2 TEXT,
    contact_info_mobile TEXT,
    contact_info_fax TEXT,
    contact_info_email TEXT,
    contact_info_mailbox_number TEXT,
    contact_info_agree_to_recieve_email BOOLEAN,
    birthday  INTEGER,
    surename TEXT NOT NULL,
    first_name TEXT NOT NULL,
    family_status_id INTEGER NOT NULL,
    FOREIGN KEY (family_status_id) REFERENCES family_statuses(id) ON DELETE CASCADE
);

CREATE table family_statuses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE patients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    local_id TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE work_periods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_id INTEGER NOT NULL,
    start DATETIME,
    end DATETIME,
    reason_id INTEGER,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (reason_id) REFERENCES end_work_period_reasons(id) ON DELETE CASCADE
);

-- Views

CREATE VIEW patients_with_emergency_details AS
SELECT  p.id,
        p.last_name,
        p.first_name,
        p.email,
        p.created_at AS patient_created_at,
        p.updated_at AS patient_updated_at,
        ed.bound_to_bed,
        ed.has_chair_wheel,
        ed.updated_at AS detail_updated_at,
        ed.last_updated_method,
        ec.last_name AS contact_last_name,
        ec.first_name AS contact_first_name,
        ec.relationship AS contact_relationship,
        ec.phone_number AS contact_phone_number,
        ec.email AS contact_email
FROM patients p
LEFT JOIN
        emergency_details ed ON p.id = ed.patient_id
LEFT JOIN 
        emergency_contacts ec ON ed.id = ec.emergency_details_id;