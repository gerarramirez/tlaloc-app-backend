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