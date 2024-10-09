-- Insert companies
INSERT INTO companies (employee_sync_ftp_address, employee_sync_ftp_input_dir, employee_sync_ftp_output_dir, employee_sync_ftp_password, employee_sync_ftp_user, name) VALUES
('62.214.45.141', 'in-hilan-employees', 'out-hilan-employees', 'admin', 'user', 'matav');

-- Insert Emergency Contacts
INSERT INTO emergency_contacts (emergency_details_id, last_name, first_name, relationship, phone_number, email) VALUES
(1, 'Smith', 'John', 'Brother', '555-1234', 'john.smith@example.com'),
(1, 'Doe', 'Jane', 'Friend', '555-5678', 'jane.doe@example.com'),
(2, 'Brown', 'Emily', 'Sister', '555-8765', 'emily.brown@example.com'),
(3, 'Johnson', 'Michael', 'Father', '555-4321', 'michael.johnson@example.com');

-- Insert employees
INSERT INTO employees (contact_info_phone_number, contact_info_phone_number2, contact_info_mobile, contact_info_fax, contact_info_email, contact_info_mailbox_number, contact_info_agree_to_recieve_email, birthday, surename, first_name, family_status_id) VALUES
('0537081811', NULL, '0537081426', NULL, 'aaroncohendev@gmail.com', NULL, 1, 123, 'Cohen', 'Aaron', 2),
('0537081426', NULL, '0537081427', NULL, 'courtneycohen58@gmail.com', NULL, 1, 568, 'Cohen', 'Courtney', 2);

-- Insert Emergency Details
INSERT INTO emergency_details (patient_id, created_at, updated_at, bound_to_bed, has_chair_wheel, last_updated_method) VALUES
(1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, FALSE, TRUE, 'manual'),
(2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, TRUE, FALSE, 'file'),
(3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, FALSE, FALSE, 'manual');

-- Insert family statues
INSERT INTO family_statuses (name) VALUES
('single'),
('married'),
('divorced'),
('widdow');

-- Insert patients
INSERT INTO patients (email, local_id, last_name, first_name, created_at, updated_at) VALUES
('alice@example.com', '341077656', 'Rowan', 'Alice', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('bob@example.com', '741077656', 'Sitar', 'Bob', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('charlie@example.com', '875670987', 'Powel', 'Charlie', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
