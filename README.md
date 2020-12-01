Project Goals: For practicing system design

### High Level Overview

*Refer from* [Design Pastebin.com](Design Pastebin.com)

#### Basic functions （Ask interviewee）

1. text -> short link 

   expire or not, expire time, delete expires

2. short link -> text

3. count url visits

#### Resource Size

1. text size
2. short link size
3. expiration length in minutes
4. created time size
5. text path

#### Expected Use

1. read/write op/s

   calculate storage

#### Database Selection and Design

1. store short link relations with text （RDBMS：use mysql in this project）
2. store path with text (OBJECT STORE)

