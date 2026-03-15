-- Races
INSERT INTO races (id, slug, name) VALUES
  (1, 'human',    'Human'),
  (2, 'halfling', 'Halfling'),
  (3, 'half-orc', 'Half-Orc');

-- ─── Human ───────────────────────────────────────────────
-- Pattern: first_name + last_name (both required)
INSERT INTO name_patterns (id, race_id, "order", component_type, required) VALUES
  (1, 1, 1, 'first_name', 1),
  (2, 1, 2, 'last_name',  1);

INSERT INTO name_components (race_id, component_type, gender, value) VALUES
  -- male first names
  (1, 'first_name', 'm', 'John'),
  (1, 'first_name', 'm', 'Peter'),
  (1, 'first_name', 'm', 'Diego'),
  (1, 'first_name', 'm', 'Marcus'),
  (1, 'first_name', 'm', 'Roland'),
  -- female first names
  (1, 'first_name', 'f', 'Elena'),
  (1, 'first_name', 'f', 'Sara'),
  (1, 'first_name', 'f', 'Marta'),
  (1, 'first_name', 'f', 'Clara'),
  (1, 'first_name', 'f', 'Diana'),
  -- last names (gender neutral)
  (1, 'last_name', NULL, 'Johnson'),
  (1, 'last_name', NULL, 'Smith'),
  (1, 'last_name', NULL, 'Martinez'),
  (1, 'last_name', NULL, 'Rivera'),
  (1, 'last_name', NULL, 'Stone');

-- ─── Halfling ─────────────────────────────────────────────
-- Pattern: first_name + composite_nickname (both required)
INSERT INTO name_patterns (id, race_id, "order", component_type, required) VALUES
  (3, 2, 1, 'first_name',          1),
  (4, 2, 2, 'composite_nickname',  1);

INSERT INTO name_components (race_id, component_type, gender, value) VALUES
  (2, 'first_name', 'm', 'Fosco'),
  (2, 'first_name', 'm', 'Bilbo'),
  (2, 'first_name', 'm', 'Drogo'),
  (2, 'first_name', 'f', 'Rosie'),
  (2, 'first_name', 'f', 'Daisy'),
  (2, 'first_name', 'f', 'Pearl');

INSERT INTO composite_parts (race_id, position, category, value) VALUES
  (2, 'first',  'noun',      'Hot'),
  (2, 'first',  'noun',      'Barrel'),
  (2, 'first',  'noun',      'Pipe'),
  (2, 'first',  'noun',      'Copper'),
  (2, 'second', 'adjective', 'Pot'),
  (2, 'second', 'adjective', 'Bottom'),
  (2, 'second', 'adjective', 'Smoke'),
  (2, 'second', 'adjective', 'Weed');

-- ─── Half-Orc ─────────────────────────────────────────────
-- Pattern: first_name only (no last name)
INSERT INTO name_patterns (id, race_id, "order", component_type, required) VALUES
  (5, 3, 1, 'first_name', 1);

INSERT INTO name_components (race_id, component_type, gender, value) VALUES
  (3, 'first_name', 'm', 'Gorth'),
  (3, 'first_name', 'm', 'Morg'),
  (3, 'first_name', 'm', 'Drak'),
  (3, 'first_name', 'm', 'Varg'),
  (3, 'first_name', 'f', 'Yrsa'),
  (3, 'first_name', 'f', 'Breka'),
  (3, 'first_name', 'f', 'Vorka');