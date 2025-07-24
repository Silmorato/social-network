-- Users table (optional if you don't manage user entities)
CREATE TABLE IF NOT EXISTS user (
                                     id TEXT PRIMARY KEY,
                                     name TEXT
);

-- Tweets table
CREATE TABLE IF NOT EXISTS tweet (
                                      id UUID PRIMARY KEY,
                                      user_id TEXT NOT NULL,
                                      content TEXT NOT NULL,
                                      created_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

-- Follows table
CREATE TABLE IF NOT EXISTS follows (
                                       follower_id TEXT NOT NULL,
                                       following_id TEXT NOT NULL,
                                       PRIMARY KEY (follower_id, following_id),
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
    );

-- Suggested indexes to improve query performance
CREATE INDEX IF NOT EXISTS idx_tweets_user_id ON tweets(user_id);
CREATE INDEX IF NOT EXISTS idx_follows_follower_id ON follows(follower_id);
CREATE INDEX IF NOT EXISTS idx_follows_following_id ON follows(following_id);