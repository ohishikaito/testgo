CREATE DATABASE IF NOT EXISTS testgo_development;
USE testgo_development;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  uid VARCHAR(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  first_name varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (uid)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table article
--

-- NOTE: 動かない
-- LOCK TABLES users WRITE;
-- /*!40000 ALTER TABLE article DISABLE KEYS */;
-- INSERT INTO users VALUES (1,'jun','makoto','2017-05-18 13:50:19','2017-05-18 13:50:19'),(2,'ohishi','kaito','2017-05-18 13:50:19','2017-05-18 13:50:19')
-- /*!40000 ALTER TABLE article ENABLE KEYS */;
-- UNLOCK TABLES;
