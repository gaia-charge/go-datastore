-- ^([\w\s]+),(.+),(.+),(.+),(.+),(.+)
-- INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('$1', '$2', $3, $4, $5, $6);

DELETE FROM htb_tariffs;
-- Paste csv below
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM139', 'CHF', NULL, NULL, 0.2, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM140', 'CHF', NULL, NULL, 0.3, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM215', 'CHF', NULL, NULL, 0.39, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM141', 'CHF', NULL, NULL, 0.5, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM216', 'CHF', NULL, NULL, 0.5571, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM217', 'CHF', NULL, NULL, 0.6035, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM218', 'CHF', NULL, NULL, 0.6500, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM219', 'CHF', NULL, NULL, 0.6964, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM220', 'CHF', NULL, NULL, 0.7428, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM221', 'CHF', NULL, NULL, 0.7892, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM222', 'CHF', NULL, NULL, 0.8357, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_100 CHF', 'CHF', 0.0154, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_250 CHF', 'CHF', 0.0387, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_500 CHF', 'CHF', 0.0772, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_800 CHF', 'CHF', 0.1235, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1000 CHF', 'CHF', 0.1543, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1200 CHF', 'CHF', 0.1852, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM198', 'CZK', NULL, NULL, 4.876, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM199', 'CZK', NULL, NULL, 5.7025, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM159', 'DKK', NULL, NULL, 2.672, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM208', 'DKK', NULL, NULL, 3.12, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM165', 'DKK', NULL, NULL, 3.568, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM209', 'DKK', NULL, NULL, 4.008, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM210', 'DKK', NULL, NULL, 4.456, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM160', 'DKK', NULL, NULL, 6.236, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM212', 'DKK', NULL, NULL, 7.127, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM213', 'DKK', NULL, NULL, 8.081, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM155', 'HUF', NULL, NULL, 82.68, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM156', 'HUF', NULL, NULL, 98.429, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM157', 'HUF', NULL, NULL, 129.921, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM158', 'HUF', NULL, NULL, 161.42, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM130', 'HUF', 26.25, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM161', 'GBP', NULL, NULL, 0.25, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM162', 'GBP', NULL, NULL, 0.275, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM163', 'GBP', NULL, NULL, 0.2917, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM164', 'GBP', NULL, NULL, 0.3333, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM187', 'GBP', NULL, NULL, 0.375, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM188', 'GBP', NULL, NULL, 0.45, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM189', 'GBP', NULL, NULL, 0.5250, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM168', 'RON', NULL, NULL, 0.9375, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM166', 'RON', NULL, NULL, 1.375, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM191', 'RON', NULL, NULL, 1.4538, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM249', 'RON', NULL, NULL, 1.5882, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM167', 'RON', NULL, NULL, 1.6875, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM192', 'RON', NULL, NULL, 1.874, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM250', 'RON', NULL, NULL, 2.0084, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM193', 'SEK', NULL, NULL, 2.36, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM242', 'SEK', NULL, NULL, 2.4, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM243', 'SEK', NULL, NULL, 3.2, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM194', 'SEK', NULL, NULL, 3.52, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM244', 'SEK', NULL, NULL, 4.00, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM240', 'SEK', NULL, NULL, 4.1600, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM245', 'SEK', NULL, NULL, 4.8, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM241', 'SEK', NULL, NULL, 5.3600, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM000', 'EUR', NULL, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM113', 'EUR', NULL, NULL, NULL, 2.94);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM114', 'EUR', NULL, NULL, NULL, 3.78);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM115', 'EUR', NULL, NULL, NULL, 6.3);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM116', 'EUR', NULL, NULL, NULL, 10.5);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM117', 'EUR', NULL, NULL, NULL, 12.61);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM118', 'EUR', NULL, NULL, NULL, 16.81);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM119', 'EUR', NULL, NULL, NULL, 25.21);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM132', 'EUR', NULL, NULL, 0.2, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM200', 'EUR', NULL, NULL, 0.2, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM186', 'EUR', NULL, NULL, 0.2083, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM202', 'EUR', NULL, NULL, 0.2097, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM251', 'EUR', NULL, NULL, 0.2364, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM203', 'EUR', NULL, NULL, 0.25, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM252', 'EUR', NULL, NULL, 0.2818, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM246', 'EUR', NULL, NULL, 0.2917, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM204', 'EUR', NULL, NULL, 0.2975, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM133', 'EUR', NULL, NULL, 0.3, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_COM133', 'EUR', NULL, NULL, 0.3, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM201', 'EUR', NULL, NULL, 0.3250, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM248', 'EUR', NULL, NULL, 0.3333, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM195', 'EUR', NULL, NULL, 0.3361, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM205', 'EUR', NULL, NULL, 0.3636, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM134', 'EUR', NULL, NULL, 0.4, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM254', 'EUR', NULL, NULL, 0.4034, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM247', 'EUR', NULL, NULL, 0.4167, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM144', 'EUR', NULL, NULL, 0.44, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM196', 'EUR', NULL, NULL, 0.4622, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC102', 'EUR', NULL, NULL, 0.575, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC105', 'EUR', NULL, NULL, 0.4900, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM256', 'EUR', NULL, NULL, 0.4917, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM228', 'EUR', NULL, NULL, 0.5125, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC104', 'EUR', NULL, NULL, 0.5400, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM145', 'EUR', NULL, NULL, 0.55, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM197', 'EUR', NULL, NULL, 0.5798, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC106', 'EUR', NULL, NULL, 0.6, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC109', 'EUR', NULL, NULL, 0.61, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC107', 'EUR', NULL, NULL, 0.65, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM253', 'EUR', NULL, NULL, 0.6555, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC101', 'EUR', NULL, NULL, 0.675, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC103', 'EUR', NULL, NULL, 0.59, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM169', 'EUR', NULL, NULL, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC108', 'EUR', NULL, NULL, 0.7, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM229', 'EUR', NULL, NULL, 0.74, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('BPC100', 'EUR', NULL, NULL, 0.775, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM146', 'EUR', NULL, NULL, 0.79, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM214', 'EUR', NULL, NULL, 0.9224, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM223', 'EUR', NULL, NULL, 1.01, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM224', 'EUR', NULL, NULL, 1.11, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM170', 'EUR', 0.10, 720, 0.33, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM171', 'EUR', 0.10, 240, 0.40, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM257', 'EUR', 0.4132, 40, 0.4132, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM172', 'EUR', 0.10, 240, 0.44, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM173', 'EUR', 0.15, 120, 0.44, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM258', 'EUR', 0.4132, 40, 0.4628, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM255', 'EUR', 0.1, 240, 0.5125, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM174', 'EUR', 0.15, 120, 0.55, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM175', 'EUR', 0.20, 60, 0.55, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM176', 'EUR', 0.20, 60, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM177', 'EUR', 0.30, 45, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM225', 'EUR', 0.0875, 180, 0.75, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM178', 'EUR', 0.30, 45, 0.79, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM226', 'EUR', 0.0875, 180, 0.85, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM227', 'EUR', 0.0875, 180, 0.95, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM179', 'EUR', NULL, NULL, 0.2389, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM180', 'EUR', NULL, NULL, 0.3019, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM181', 'EUR', 0.1161, 180, 0.3567, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM182', 'EUR', 0.1161, 180, 0.4645, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM190', 'EUR', 0.1161, 180, 0.6304, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM183', 'CHF', 0.0929, 720, 0.3714, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM184', 'CHF', 0.0929, 240, 0.4178, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM185', 'CHF', 0.1393, 120, 0.4178, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM235', 'CHF', 0.125, 240, 0.4375, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM236', 'CHF', 0.125, 240, 0.5, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM232', 'CHF', 0.125, 240, 0.5625, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM233', 'CHF', 0.3125, 60, 0.625, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM237', 'CHF', 0.125, 240, 0.625, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM230', 'CHF', 0.125, 240, 0.68, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM231', 'CHF', 0.3125, 60, 0.68, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM234', 'CHF', 0.3125, 60, 0.73, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM238', 'CHF', 0.3125, 60, 0.78, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM239', 'CHF', 0.3125, 60, 0.83, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_100 EUR', 'EUR', 0.0166, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_250 EUR', 'EUR', 0.0417, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_350 EUR', 'EUR', 0.0583, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_500 EUR', 'EUR', 0.0833, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_650 EUR', 'EUR', 0.1083, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_800 EUR', 'EUR', 0.1333, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1000 EUR', 'EUR', 0.1667, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM142', 'EUR', 0.1875, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1200 EUR', 'EUR', 0.2, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1500 EUR', 'EUR', 0.25, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_2000 EUR', 'EUR', 0.3333, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_2500 EUR', 'EUR', 0.4167, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM211', 'EUR', 0.4583, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_3000 EUR', 'EUR', 0.5, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_4500 EUR', 'EUR', 0.75, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_8000 EUR', 'EUR', 1.3333, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_450 EUR PKG', 'EUR', 0.075, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_1500 EUR PKG', 'EUR', 0.25, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM206', 'SEK', 0.248, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_16 SEK PKG', 'SEK', 0.2666, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_20_SEK_PKG', 'SEK', 0.3334, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_24_SEK_PKG', 'SEK', 0.4, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('COM207', 'SEK', 0.416, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_32 SEK PKG', 'SEK', 0.5333, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_40_SEK_PKG', 'SEK', 0.6666, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_48_SEK_PKG', 'SEK', 0.8, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_64_SEK_PKG', 'SEK', 1.0666, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_80 SEK PKG', 'SEK', 1.3333, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_T_96_SEK_PKG', 'SEK', 1.6, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION200', 'EUR', NULL, NULL, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION200_MC', 'EUR', NULL, NULL, 0.41, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION201', 'EUR', NULL, NULL, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION202', 'CHF', NULL, NULL, 0.73, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION203', 'EUR', NULL, NULL, 0.65, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION203_MC', 'EUR', NULL, NULL, 0.4, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION204', 'GBP', NULL, NULL, 0.62, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION205', 'NOK', NULL, NULL, 6.72, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION206', 'SEK', NULL, NULL, 6.96, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION207', 'DKK', NULL, NULL, 4.96, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION208', 'EUR', NULL, NULL, 0.65, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION208_MC', 'EUR', NULL, NULL, 0.4, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION209', 'EUR', NULL, NULL, 0.64, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION211', 'HUF', NULL, NULL, 220.47, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION212', 'EUR', NULL, NULL, 0.66, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION213', 'CZK', NULL, NULL, 17.36, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION214', 'EUR', NULL, NULL, 0.64, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION215', 'PLN', NULL, NULL, 2.85, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION216', 'EUR', NULL, NULL, 0.51, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION217', 'EUR', NULL, NULL, 0.54, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION314', 'EUR', 0.64, NULL, NULL, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION317', 'EUR', NULL, NULL, 0.57, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('ION317_MC', 'EUR', NULL, NULL, 0.33, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_AC3', 'EUR', NULL, NULL, 0.2521, 0.8319);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_DC3', 'EUR', NULL, NULL, 0.2941, 0.8319);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_HPC3', 'EUR', NULL, NULL, 0.2941, 0.8319);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_HUB3', 'EUR', 0.084, 45, 0.3, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_M_AC3', 'EUR', NULL, NULL, 0.3277, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_M_HPC3', 'EUR', NULL, NULL, 0.4118, NULL);
INSERT INTO htb_tariffs (name, currency, time_price, time_min_duration, energy_price, flat_price) VALUES ('MC_AUC_CHF_HUB3', 'CHF', 0.0929, 45, 0.3, NULL);