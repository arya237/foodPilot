SELECT id, user_id, provider, identifier
FROM identities
WHERE provider = 'telegram' --and identifier == "d"