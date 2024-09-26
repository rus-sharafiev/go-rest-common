CREATE TYPE access AS ENUM ('USER', 'ADMIN');
CREATE TABLE users (
    "id" serial PRIMARY KEY,
    "email" text UNIQUE,
    "firstName" text,
    "lastName" text,
    "phone" text UNIQUE,
    "avatar" text,
    "access" access DEFAULT 'USER',
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP WITH TIME ZONE
);
CREATE TABLE passwords (
    "userId" integer REFERENCES users (id) ON DELETE CASCADE,
    "passwordHash" text
);
CREATE TABLE sessions (
    "userId" integer REFERENCES users (id) ON DELETE CASCADE,
    "fingerprint" text UNIQUE,
    "userAgent" text,
    "ip" text,
    "updatedAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE chats (
    "id" serial PRIMARY KEY,
    "users" integer [] UNIQUE,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP WITH TIME ZONE
);
CREATE TABLE messages (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "chatId" integer REFERENCES chats (id) ON DELETE CASCADE,
    "from" integer REFERENCES users (id) ON DELETE CASCADE,
    "to" integer REFERENCES users (id) ON DELETE CASCADE,
    "message" text,
    "file" text,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deliveredAt" TIMESTAMP WITH TIME ZONE,
    "readAt" TIMESTAMP WITH TIME ZONE
);