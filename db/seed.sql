-- Insert Users
INSERT INTO patients (email, local_id, last_name, first_name, created_at, updated_at) VALUES
('alice@example.com', '341077656', 'Rowan', 'Alice', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('bob@example.com', '741077656', 'Sitar', 'Bob', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('charlie@example.com', '875670987', 'Powel', 'Charlie', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert Emergency Details
INSERT INTO emergency_details (patient_id, created_at, updated_at, bound_to_bed, has_chair_wheel, last_updated_method) VALUES
(1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, FALSE, TRUE, 'manual'),
(2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, TRUE, FALSE, 'file'),
(3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, FALSE, FALSE, 'manual');

-- Insert Emergency Contacts
INSERT INTO emergency_contacts (emergency_details_id, last_name, first_name, relationship, phone_number, email) VALUES
(1, 'Smith', 'John', 'Brother', '555-1234', 'john.smith@example.com'),
(1, 'Doe', 'Jane', 'Friend', '555-5678', 'jane.doe@example.com'),
(2, 'Brown', 'Emily', 'Sister', '555-8765', 'emily.brown@example.com'),
(3, 'Johnson', 'Michael', 'Father', '555-4321', 'michael.johnson@example.com');