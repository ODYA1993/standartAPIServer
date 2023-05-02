CREATE TABLE friends
(
    id        serial,
    user   INT NOT NULL,
    friend INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (friend_id) REFERENCES users (id)
);