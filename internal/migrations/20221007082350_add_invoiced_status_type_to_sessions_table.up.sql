-- Sessions
ALTER TYPE session_status_type ADD VALUE IF NOT EXISTS 'INVOICED' AFTER 'INVALID';