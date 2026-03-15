-- ─────────────────────────────────────────────
-- RACES
-- ─────────────────────────────────────────────
INSERT INTO races (id, slug, name, parent_race_id) VALUES
  (1, 'human',      'Human',      NULL),
  (2, 'dwarf',      'Dwarf',      NULL),
  (3, 'elf',        'Elf',        NULL),
  (4, 'halfling',   'Halfling',   NULL),
  (5, 'dragonborn', 'Dragonborn', NULL),
  (6, 'gnome',      'Gnome',      NULL),
  (7, 'half-elf',   'Half-Elf',   NULL),
  (8, 'half-orc',   'Half-Orc',   NULL),
  (9, 'tiefling',   'Tiefling',   NULL);

-- ─────────────────────────────────────────────
-- NAME STYLES
-- ─────────────────────────────────────────────
INSERT INTO name_styles (id, race_id, slug) VALUES
  (1, 9, 'infernal'),
  (2, 9, 'virtue'),
  (3, 7, 'human'),
  (4, 7, 'elven');

-- ─────────────────────────────────────────────
-- NAME PATTERNS
-- ─────────────────────────────────────────────

-- Human: first_name + last_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (1, 1, NULL, 1, 'first_name', 1, 1),
  (2, 1, NULL, 2, 'last_name',  1, 1);

-- Dwarf: first_name + clan_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (3, 2, NULL, 1, 'first_name', 1, 1),
  (4, 2, NULL, 2, 'clan_name',  1, 1);

-- Elf: adult_name + family_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (5, 3, NULL, 1, 'adult_name',  1, 1),
  (6, 3, NULL, 2, 'family_name', 1, 1);

-- Halfling: first_name + last_name (composite) + optional nickname
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (7,  4, NULL, 1, 'first_name', 1, 1),
  (8,  4, NULL, 2, 'last_name',  1, 1),
  (9,  4, NULL, 3, 'nickname',   0, 1);

-- Dragonborn: clan_name (first!) + first_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (10, 5, NULL, 1, 'clan_name',  1, 1),
  (11, 5, NULL, 2, 'first_name', 1, 1);

-- Gnome: first_name + clan_name + up to 3 nicknames
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (12, 6, NULL, 1, 'first_name', 1, 1),
  (13, 6, NULL, 2, 'clan_name',  1, 1),
  (14, 6, NULL, 3, 'nickname',   0, 3);

-- Half-Elf human style: first_name + last_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (15, 7, 3, 1, 'first_name', 1, 1),
  (16, 7, 3, 2, 'last_name',  1, 1);

-- Half-Elf elven style: adult_name + family_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (17, 7, 4, 1, 'adult_name',  1, 1),
  (18, 7, 4, 2, 'family_name', 1, 1);

-- Half-Orc: first_name + optional last_name
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (19, 8, NULL, 1, 'first_name', 1, 1),
  (20, 8, NULL, 2, 'last_name',  0, 1);

-- Tiefling infernal style
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (21, 9, 1, 1, 'infernal_name', 1, 1);

-- Tiefling virtue style
INSERT INTO name_patterns (id, race_id, style_id, "order", component_type, required, max_count) VALUES
  (22, 9, 2, 1, 'virtue_name', 1, 1);

-- ─────────────────────────────────────────────
-- NAME COMPONENTS
-- ─────────────────────────────────────────────

-- Human
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (1,  1, 'first_name', 'm', 'Aldric'),
  (2,  1, 'first_name', 'm', 'Bram'),
  (3,  1, 'first_name', 'm', 'Cedric'),
  (4,  1, 'first_name', 'm', 'Dorian'),
  (5,  1, 'first_name', 'm', 'Edric'),
  (6,  1, 'first_name', 'f', 'Aelindra'),
  (7,  1, 'first_name', 'f', 'Brynn'),
  (8,  1, 'first_name', 'f', 'Calla'),
  (9,  1, 'first_name', 'f', 'Dwyn'),
  (10, 1, 'first_name', 'f', 'Elara'),
  (11, 1, 'last_name',  NULL, 'Ashford'),
  (12, 1, 'last_name',  NULL, 'Blackwood'),
  (13, 1, 'last_name',  NULL, 'Cresthill'),
  (14, 1, 'last_name',  NULL, 'Dunmore'),
  (15, 1, 'last_name',  NULL, 'Evermoor');

-- Dwarf
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (16, 2, 'first_name', 'm', 'Balin'),
  (17, 2, 'first_name', 'm', 'Dain'),
  (18, 2, 'first_name', 'm', 'Gimli'),
  (19, 2, 'first_name', 'm', 'Thorin'),
  (20, 2, 'first_name', 'm', 'Oin'),
  (21, 2, 'first_name', 'f', 'Hlin'),
  (22, 2, 'first_name', 'f', 'Kathra'),
  (23, 2, 'first_name', 'f', 'Mardred'),
  (24, 2, 'first_name', 'f', 'Tordrid'),
  (25, 2, 'first_name', 'f', 'Vistra'),
  (26, 2, 'clan_name',  NULL, 'Balderk'),
  (27, 2, 'clan_name',  NULL, 'Dankil'),
  (28, 2, 'clan_name',  NULL, 'Gorunn'),
  (29, 2, 'clan_name',  NULL, 'Loderr'),
  (30, 2, 'clan_name',  NULL, 'Lutgehr');

-- Elf
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (31, 3, 'adult_name',  'm', 'Adran'),
  (32, 3, 'adult_name',  'm', 'Aramil'),
  (33, 3, 'adult_name',  'm', 'Carric'),
  (34, 3, 'adult_name',  'm', 'Erevan'),
  (35, 3, 'adult_name',  'm', 'Galinndan'),
  (36, 3, 'adult_name',  'f', 'Adrie'),
  (37, 3, 'adult_name',  'f', 'Althaea'),
  (38, 3, 'adult_name',  'f', 'Anastrianna'),
  (39, 3, 'adult_name',  'f', 'Andraste'),
  (40, 3, 'adult_name',  'f', 'Antinua'),
  (41, 3, 'family_name', NULL, 'Amakiir'),
  (42, 3, 'family_name', NULL, 'Amastacia'),
  (43, 3, 'family_name', NULL, 'Galanodel'),
  (44, 3, 'family_name', NULL, 'Holimion'),
  (45, 3, 'family_name', NULL, 'Liadon');

-- Halfling
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (46, 4, 'first_name', 'm', 'Alton'),
  (47, 4, 'first_name', 'm', 'Ander'),
  (48, 4, 'first_name', 'm', 'Cade'),
  (49, 4, 'first_name', 'm', 'Corrin'),
  (50, 4, 'first_name', 'm', 'Eldon'),
  (51, 4, 'first_name', 'f', 'Andry'),
  (52, 4, 'first_name', 'f', 'Bree'),
  (53, 4, 'first_name', 'f', 'Callie'),
  (54, 4, 'first_name', 'f', 'Cora'),
  (55, 4, 'first_name', 'f', 'Euphemia'),
  (56, 4, 'nickname',   NULL, 'Three-fingers'),
  (57, 4, 'nickname',   NULL, 'Bitterleaf'),
  (58, 4, 'nickname',   NULL, 'the Lucky'),
  (59, 4, 'nickname',   NULL, 'Quickfoot'),
  (60, 4, 'nickname',   NULL, 'Dusty');

-- Dragonborn
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (61, 5, 'first_name', 'm', 'Arjhan'),
  (62, 5, 'first_name', 'm', 'Balasar'),
  (63, 5, 'first_name', 'm', 'Bharash'),
  (64, 5, 'first_name', 'm', 'Donaar'),
  (65, 5, 'first_name', 'm', 'Ghesh'),
  (66, 5, 'first_name', 'f', 'Akra'),
  (67, 5, 'first_name', 'f', 'Biri'),
  (68, 5, 'first_name', 'f', 'Dort'),
  (69, 5, 'first_name', 'f', 'Farideh'),
  (70, 5, 'first_name', 'f', 'Harann'),
  (71, 5, 'clan_name',  NULL, 'Clethtinthiallor'),
  (72, 5, 'clan_name',  NULL, 'Daardendrian'),
  (73, 5, 'clan_name',  NULL, 'Delmirev'),
  (74, 5, 'clan_name',  NULL, 'Drachedandion'),
  (75, 5, 'clan_name',  NULL, 'Fenkenkabradon');

-- Gnome
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (76,  6, 'first_name', 'm', 'Alston'),
  (77,  6, 'first_name', 'm', 'Alvyn'),
  (78,  6, 'first_name', 'm', 'Boddynock'),
  (79,  6, 'first_name', 'm', 'Brocc'),
  (80,  6, 'first_name', 'm', 'Burgell'),
  (81,  6, 'first_name', 'f', 'Bimpnottin'),
  (82,  6, 'first_name', 'f', 'Breena'),
  (83,  6, 'first_name', 'f', 'Caramip'),
  (84,  6, 'first_name', 'f', 'Carlin'),
  (85,  6, 'first_name', 'f', 'Donella'),
  (86,  6, 'clan_name',  NULL, 'Beren'),
  (87,  6, 'clan_name',  NULL, 'Daergel'),
  (88,  6, 'clan_name',  NULL, 'Folkor'),
  (89,  6, 'clan_name',  NULL, 'Garrick'),
  (90,  6, 'clan_name',  NULL, 'Nackle'),
  (91,  6, 'nickname',   NULL, 'Aleslosh'),
  (92,  6, 'nickname',   NULL, 'Ashhearth'),
  (93,  6, 'nickname',   NULL, 'Badger'),
  (94,  6, 'nickname',   NULL, 'Cloak'),
  (95,  6, 'nickname',   NULL, 'Doublelock'),
  (96,  6, 'nickname',   NULL, 'Filchbatter'),
  (97,  6, 'nickname',   NULL, 'Fnipper'),
  (98,  6, 'nickname',   NULL, 'Ku'),
  (99,  6, 'nickname',   NULL, 'Nim'),
  (100, 6, 'nickname',   NULL, 'Oneshoe');

-- Half-Elf
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (101, 7, 'first_name',  'm', 'Aramil'),
  (102, 7, 'first_name',  'm', 'Berrian'),
  (103, 7, 'first_name',  'm', 'Dayereth'),
  (104, 7, 'first_name',  'f', 'Caelynn'),
  (105, 7, 'first_name',  'f', 'Ielenia'),
  (106, 7, 'last_name',   NULL, 'Brightwood'),
  (107, 7, 'last_name',   NULL, 'Moonwhisper'),
  (108, 7, 'last_name',   NULL, 'Silverwind'),
  (109, 7, 'adult_name',  'm', 'Arannis'),
  (110, 7, 'adult_name',  'f', 'Sariel'),
  (111, 7, 'family_name', NULL, 'Nailo'),
  (112, 7, 'family_name', NULL, 'Siannodel');

-- Half-Orc
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (113, 8, 'first_name', 'm', 'Dench'),
  (114, 8, 'first_name', 'm', 'Feng'),
  (115, 8, 'first_name', 'm', 'Gell'),
  (116, 8, 'first_name', 'm', 'Morg'),
  (117, 8, 'first_name', 'm', 'Ront'),
  (118, 8, 'first_name', 'f', 'Baggi'),
  (119, 8, 'first_name', 'f', 'Emen'),
  (120, 8, 'first_name', 'f', 'Engong'),
  (121, 8, 'first_name', 'f', 'Kansif'),
  (122, 8, 'first_name', 'f', 'Sutha'),
  (123, 8, 'last_name',  NULL, 'Barton'),
  (124, 8, 'last_name',  NULL, 'Cord'),
  (125, 8, 'last_name',  NULL, 'Hadley'),
  (126, 8, 'last_name',  NULL, 'Marsh'),
  (127, 8, 'last_name',  NULL, 'Stirling');

-- Tiefling
INSERT INTO name_components (id, race_id, component_type, gender, value) VALUES
  (128, 9, 'infernal_name', 'm', 'Akmenos'),
  (129, 9, 'infernal_name', 'm', 'Amnon'),
  (130, 9, 'infernal_name', 'm', 'Barakas'),
  (131, 9, 'infernal_name', 'm', 'Damakos'),
  (132, 9, 'infernal_name', 'm', 'Ekemon'),
  (133, 9, 'infernal_name', 'f', 'Akta'),
  (134, 9, 'infernal_name', 'f', 'Anakis'),
  (135, 9, 'infernal_name', 'f', 'Bryseis'),
  (136, 9, 'infernal_name', 'f', 'Criella'),
  (137, 9, 'infernal_name', 'f', 'Damaia'),
  (138, 9, 'virtue_name',   NULL, 'Art'),
  (139, 9, 'virtue_name',   NULL, 'Carrion'),
  (140, 9, 'virtue_name',   NULL, 'Chant'),
  (141, 9, 'virtue_name',   NULL, 'Despair'),
  (142, 9, 'virtue_name',   NULL, 'Excellence'),
  (143, 9, 'virtue_name',   NULL, 'Fear'),
  (144, 9, 'virtue_name',   NULL, 'Glory'),
  (145, 9, 'virtue_name',   NULL, 'Hope'),
  (146, 9, 'virtue_name',   NULL, 'Ideal'),
  (147, 9, 'virtue_name',   NULL, 'Torment');

-- ─────────────────────────────────────────────
-- COMPOSITE PARTS (Halfling last names)
-- "Brush" + "gather" → "Brushgather"
-- ─────────────────────────────────────────────
INSERT INTO composite_parts (id, race_id, position, category, value) VALUES
  (1,  4, 'first',  'noun', 'Brush'),
  (2,  4, 'first',  'noun', 'Good'),
  (3,  4, 'first',  'noun', 'High'),
  (4,  4, 'first',  'noun', 'Tall'),
  (5,  4, 'first',  'noun', 'Green'),
  (6,  4, 'first',  'noun', 'Copper'),
  (7,  4, 'first',  'noun', 'Bright'),
  (8,  4, 'first',  'noun', 'Merry'),
  (9,  4, 'second', 'verb', 'gather'),
  (10, 4, 'second', 'noun', 'barrel'),
  (11, 4, 'second', 'noun', 'hill'),
  (12, 4, 'second', 'noun', 'tree'),
  (13, 4, 'second', 'noun', 'leaf'),
  (14, 4, 'second', 'noun', 'foot'),
  (15, 4, 'second', 'noun', 'bottom'),
  (16, 4, 'second', 'noun', 'hollow');