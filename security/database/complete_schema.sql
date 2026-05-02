-- ============================================
-- Tlaloc Security Service - Database Schema Completo
-- ============================================

-- Set search path for correct schema
SET search_path TO tlaloc_security_user, public;

-- ============================================
-- Drop existing tables (clean recreation)
-- ============================================
DROP TABLE IF EXISTS tlaloc_security_user.user_sessions CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.refresh_tokens CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.auth_challenges CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.user_role_assignments CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.users CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.role_privileges CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.roles CASCADE;
DROP TABLE IF EXISTS tlaloc_security_user.privileges CASCADE;

-- ============================================
-- Privileges Table
-- ============================================
CREATE TABLE tlaloc_security_user.privileges (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    resource VARCHAR(100) NOT NULL, -- e.g., 'users', 'transactions', 'budgets'
    action VARCHAR(50) NOT NULL,   -- e.g., 'create', 'read', 'update', 'delete'
    is_system BOOLEAN NOT NULL DEFAULT false, -- System privileges cannot be deleted
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_privileges_name ON tlaloc_security_user.privileges(name);
CREATE INDEX idx_privileges_resource ON tlaloc_security_user.privileges(resource);
CREATE INDEX idx_privileges_action ON tlaloc_security_user.privileges(action);
CREATE UNIQUE INDEX idx_privileges_resource_action ON tlaloc_security_user.privileges(resource, action);

-- ============================================
-- Roles Table
-- ============================================
CREATE TABLE tlaloc_security_user.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    is_system BOOLEAN NOT NULL DEFAULT false, -- System roles cannot be deleted
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX idx_roles_name ON tlaloc_security_user.roles(name);
CREATE INDEX idx_roles_is_active ON tlaloc_security_user.roles(is_active);
CREATE INDEX idx_roles_deleted_at ON tlaloc_security_user.roles(deleted_at);

-- ============================================
-- Role Privileges Junction Table
-- ============================================
CREATE TABLE tlaloc_security_user.role_privileges (
    id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL,
    privilege_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES tlaloc_security_user.roles(id) ON DELETE CASCADE,
    FOREIGN KEY (privilege_id) REFERENCES tlaloc_security_user.privileges(id) ON DELETE CASCADE,
    UNIQUE(role_id, privilege_id)
);

-- Create indexes
CREATE INDEX idx_role_privileges_role_id ON tlaloc_security_user.role_privileges(role_id);
CREATE INDEX idx_role_privileges_privilege_id ON tlaloc_security_user.role_privileges(privilege_id);

-- ============================================
-- Users Table (Updated without hardcoded role)
-- ============================================
CREATE TABLE tlaloc_security_user.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX idx_users_email ON tlaloc_security_user.users(email);
CREATE INDEX idx_users_is_active ON tlaloc_security_user.users(is_active);
CREATE INDEX idx_users_deleted_at ON tlaloc_security_user.users(deleted_at);

-- ============================================
-- User Role Assignments Junction Table
-- ============================================
CREATE TABLE tlaloc_security_user.user_role_assignments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    assigned_by INTEGER, -- User who assigned this role
    is_active BOOLEAN NOT NULL DEFAULT true,
    expires_at TIMESTAMP WITH TIME ZONE, -- Optional role expiration
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES tlaloc_security_user.users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES tlaloc_security_user.roles(id) ON DELETE CASCADE,
    FOREIGN KEY (assigned_by) REFERENCES tlaloc_security_user.users(id) ON DELETE SET NULL,
    UNIQUE(user_id, role_id)
);

-- Create indexes
CREATE INDEX idx_user_role_assignments_user_id ON tlaloc_security_user.user_role_assignments(user_id);
CREATE INDEX idx_user_role_assignments_role_id ON tlaloc_security_user.user_role_assignments(role_id);
CREATE INDEX idx_user_role_assignments_is_active ON tlaloc_security_user.user_role_assignments(is_active);
CREATE INDEX idx_user_role_assignments_expires_at ON tlaloc_security_user.user_role_assignments(expires_at);

-- ============================================
-- Auth Challenges Table (NEW for Challenge-Response)
-- ============================================
CREATE TABLE tlaloc_security_user.auth_challenges (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    challenge VARCHAR(64) NOT NULL,
    nonce VARCHAR(32) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for challenges
CREATE INDEX idx_auth_challenges_email ON tlaloc_security_user.auth_challenges(email);
CREATE INDEX idx_auth_challenges_challenge ON tlaloc_security_user.auth_challenges(challenge);
CREATE INDEX idx_auth_challenges_expires_at ON tlaloc_security_user.auth_challenges(expires_at);
CREATE INDEX idx_auth_challenges_is_used ON tlaloc_security_user.auth_challenges(is_used);
CREATE UNIQUE INDEX idx_auth_challenges_email_unique 
    ON tlaloc_security_user.auth_challenges(email) WHERE is_used = false;

-- ============================================
-- Refresh Tokens Table
-- ============================================
CREATE TABLE tlaloc_security_user.refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token VARCHAR(500) UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES tlaloc_security_user.users(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_refresh_tokens_user_id ON tlaloc_security_user.refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON tlaloc_security_user.refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON tlaloc_security_user.refresh_tokens(expires_at);
CREATE INDEX idx_refresh_tokens_is_active ON tlaloc_security_user.refresh_tokens(is_active);
CREATE INDEX idx_refresh_tokens_deleted_at ON tlaloc_security_user.refresh_tokens(deleted_at);

-- ============================================
-- User Sessions Table
-- ============================================
CREATE TABLE tlaloc_security_user.user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token_jti VARCHAR(100) UNIQUE NOT NULL,
    ip_address INET,
    user_agent TEXT,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES tlaloc_security_user.users(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_user_sessions_user_id ON tlaloc_security_user.user_sessions(user_id);
CREATE INDEX idx_user_sessions_token_jti ON tlaloc_security_user.user_sessions(token_jti);
CREATE INDEX idx_user_sessions_expires_at ON tlaloc_security_user.user_sessions(expires_at);
CREATE INDEX idx_user_sessions_is_active ON tlaloc_security_user.user_sessions(is_active);
CREATE INDEX idx_user_sessions_deleted_at ON tlaloc_security_user.user_sessions(deleted_at);

-- ============================================
-- Update timestamp function
-- ============================================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON tlaloc_security_user.users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_roles_updated_at BEFORE UPDATE ON tlaloc_security_user.roles
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_privileges_updated_at BEFORE UPDATE ON tlaloc_security_user.privileges
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_refresh_tokens_updated_at BEFORE UPDATE ON tlaloc_security_user.refresh_tokens
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_sessions_updated_at BEFORE UPDATE ON tlaloc_security_user.user_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_role_assignments_updated_at BEFORE UPDATE ON tlaloc_security_user.user_role_assignments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_auth_challenges_updated_at BEFORE UPDATE ON tlaloc_security_user.auth_challenges
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();


-- ============================================
-- Advanced Functions for Challenge-Response
-- ============================================

-- Function to cleanup expired challenges
CREATE OR REPLACE FUNCTION cleanup_expired_challenges()
RETURNS INTEGER AS $$
DECLARE
    deleted_count INTEGER;
BEGIN
    DELETE FROM tlaloc_security_user.auth_challenges 
    WHERE expires_at < CURRENT_TIMESTAMP 
    OR is_used = true;
    
    GET DIAGNOSTICS deleted_count = ROW_COUNT;
    RETURN deleted_count;
END;
$$ LANGUAGE plpgsql;

-- Function to validate and use challenge
CREATE OR REPLACE FUNCTION validate_and_use_challenge(
    p_email VARCHAR(255),
    p_challenge VARCHAR(64),
    p_nonce VARCHAR(32),
    p_response_hash VARCHAR(64)
) RETURNS BOOLEAN AS $$
DECLARE
    challenge_record RECORD;
    user_record RECORD;
    is_valid BOOLEAN := false;
BEGIN
    -- 1. Find the challenge
    SELECT * INTO challenge_record
    FROM tlaloc_security_user.auth_challenges 
    WHERE email = p_email 
    AND challenge = p_challenge 
    AND nonce = p_nonce 
    AND is_used = false 
    AND expires_at > CURRENT_TIMESTAMP
    LIMIT 1;
    
    -- If not found, return false
    IF NOT FOUND THEN
        RETURN false;
    END IF;
    
    -- 2. Get the user
    SELECT * INTO user_record
    FROM tlaloc_security_user.users 
    WHERE email = p_email 
    AND is_active = true 
    AND deleted_at IS NULL;
    
    -- If user not found or inactive, mark challenge as used for security
    IF NOT FOUND THEN
        UPDATE tlaloc_security_user.auth_challenges 
        SET is_used = true,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = challenge_record.id;
        RETURN false;
    END IF;
    
    -- 3. Return true (challenge is valid for Go to verify hash)
    RETURN true;
END;
$$ LANGUAGE plpgsql;

-- Function to mark challenge as used
CREATE OR REPLACE FUNCTION mark_challenge_used(
    p_email VARCHAR(255),
    p_challenge VARCHAR(64)
) RETURNS BOOLEAN AS $$
BEGIN
    UPDATE tlaloc_security_user.auth_challenges 
    SET is_used = true,
        updated_at = CURRENT_TIMESTAMP
    WHERE email = p_email 
    AND challenge = p_challenge 
    AND is_used = false;
    
    RETURN true;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- Fix para las VISTAS (CORRECCIÓN)
-- ============================================

-- Reemplazar las vistas existentes con estas correcciones:

-- ============================================
-- View for user roles and their privileges (CORREGIDO)
-- ============================================
CREATE OR REPLACE VIEW tlaloc_security_user.user_role_privileges AS
SELECT 
    u.id as user_id,
    u.email,
    u.first_name,
    u.last_name,
    r.id as role_id,
    r.name as role_name,
    r.description as role_description,
    p.id as privilege_id,
    p.name as privilege_name,
    p.description as privilege_description,
    p.resource,
    p.action,
    ura.is_active as role_assignment_active,
    ura.expires_at as role_assignment_expires_at
FROM tlaloc_security_user.users u
JOIN tlaloc_security_user.user_role_assignments ura ON u.id = ura.user_id
JOIN tlaloc_security_user.roles r ON ura.role_id = r.id
JOIN tlaloc_security_user.role_privileges rp ON r.id = rp.role_id
JOIN tlaloc_security_user.privileges p ON rp.privilege_id = p.id
WHERE u.deleted_at IS NULL 
    AND u.is_active = true
    AND ura.is_active = true
    AND (ura.expires_at IS NULL OR ura.expires_at > CURRENT_TIMESTAMP)
    AND r.deleted_at IS NULL
    AND r.is_active = true
    AND p.deleted_at IS NULL;

-- ============================================
-- View for active users with their tokens and roles (CORREGIDO)
-- ============================================
CREATE OR REPLACE VIEW tlaloc_security_user.active_users_tokens AS
SELECT 
    u.id,
    u.email,
    u.first_name,
    u.last_name,
    u.is_active,
    u.created_at,
    u.updated_at,
    COUNT(rt.id) as active_refresh_tokens,
    array_agg(DISTINCT r.name) FILTER (WHERE r.name IS NOT NULL) as roles
FROM tlaloc_security_user.users u
LEFT JOIN tlaloc_security_user.refresh_tokens rt ON u.id = rt.user_id AND rt.is_active = true AND rt.expires_at > CURRENT_TIMESTAMP
LEFT JOIN tlaloc_security_user.user_role_assignments ura ON u.id = ura.user_id AND ura.is_active = true 
    AND (ura.expires_at IS NULL OR ura.expires_at > CURRENT_TIMESTAMP)
LEFT JOIN tlaloc_security_user.roles r ON ura.role_id = r.id AND r.is_active = true AND r.deleted_at IS NULL
WHERE u.is_active = true AND u.deleted_at IS NULL
GROUP BY u.id, u.email, u.first_name, u.last_name, u.is_active, u.created_at, u.updated_at;

-- ============================================
-- View for user sessions (CORREGIDO)
-- ============================================
CREATE OR REPLACE VIEW tlaloc_security_user.user_active_sessions AS
SELECT 
    us.id,
    us.user_id,
    u.email,
    array_agg(DISTINCT r.name) FILTER (WHERE r.name IS NOT NULL) as user_roles,
    us.token_jti,
    us.ip_address,
    us.user_agent,
    us.expires_at,
    us.created_at
FROM tlaloc_security_user.user_sessions us
JOIN tlaloc_security_user.users u ON us.user_id = u.id
LEFT JOIN tlaloc_security_user.user_role_assignments ura ON us.user_id = ura.user_id AND ura.is_active = true 
    AND (ura.expires_at IS NULL OR ura.expires_at > CURRENT_TIMESTAMP)
LEFT JOIN tlaloc_security_user.roles r ON ura.role_id = r.id AND r.is_active = true AND r.deleted_at IS NULL
WHERE us.is_active = true 
    AND us.expires_at > CURRENT_TIMESTAMP
    AND us.deleted_at IS NULL
    AND u.is_active = true
    AND u.deleted_at IS NULL
GROUP BY us.id, us.user_id, us.token_jti, us.ip_address, us.user_agent, us.expires_at, us.created_at;

-- ============================================
-- Mensaje de confirmación
-- ============================================
SELECT 'Database views fixed successfully!' as message;

-- ============================================
-- Comments
-- ============================================

COMMENT ON TABLE tlaloc_security_user.users IS 'Users table for authentication and authorization';
COMMENT ON TABLE tlaloc_security_user.roles IS 'Roles table for role-based access control';
COMMENT ON TABLE tlaloc_security_user.privileges IS 'Privileges table defining actions on resources';
COMMENT ON TABLE tlaloc_security_user.role_privileges IS 'Junction table linking roles and privileges';
COMMENT ON TABLE tlaloc_security_user.user_role_assignments IS 'Junction table linking users and roles with assignment tracking';
COMMENT ON TABLE tlaloc_security_user.auth_challenges IS 'Challenge-response authentication table';
COMMENT ON TABLE tlaloc_security_user.refresh_tokens IS 'Refresh tokens for JWT renewal mechanism';
COMMENT ON TABLE tlaloc_security_user.user_sessions IS 'Active user sessions for tracking and audit';

COMMENT ON COLUMN tlaloc_security_user.users.email IS 'User email address, must be unique';
COMMENT ON COLUMN tlaloc_security_user.users.password IS 'Hashed password using Argon2ID';
COMMENT ON COLUMN tlaloc_security_user.auth_challenges.challenge IS 'Random challenge string';
COMMENT ON COLUMN tlaloc_security_user.auth_challenges.nonce IS 'Random nonce for additional security';
COMMENT ON COLUMN tlaloc_security_user.auth_challenges.expires_at IS 'Challenge expiration timestamp';
COMMENT ON COLUMN tlaloc_security_user.auth_challenges.is_used IS 'Whether challenge has been used';

-- Reset search path
RESET search_path;

-- ============================================
-- Grant permissions to application user
-- ============================================

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA tlaloc_security_user TO tlaloc_security_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA tlaloc_security_user TO tlaloc_security_user;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA tlaloc_security_user TO tlaloc_security_user;

SELECT 'Complete database schema created successfully!' as message;