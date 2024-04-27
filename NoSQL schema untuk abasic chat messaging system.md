## NoSQL schema untuk abasic chat messaging system

### user

```json
{
  "_id": "user_1",
  "name": "John Doe",
  "phone_number": "+1234567890",
  "password": "hashed_password",
  "created_at": "2022-01-01T00:00:00.000Z",
  "updated_at": "2022-01-01T12:00:00.000Z"
}
```

### conversation

```json
{
  "_id": "conversation_123",
  "sender_id": "user_1",
  "receiver_id": "user_2",
  "messages": [
    {
      "_id": "message_1",
      "sender_id": "user_1",
      "receiver_id": "user_2",
      "content": "Halo, apa kabar?",
      "timestamp": "2022-01-01T08:00:00.000Z",
      "status": "sent",
      "attachment": [{
        "name": "greet.jpg",
        "type": "image",
        "url": "https://dummyimage.com/600x400/000/fff&text=Halo,+apa+kabar?"
      }]
    },
    {
      "_id": "message_2",
      "sender_id": "user_2",
      "receiver_id": "user_1",
      "content": "Halo juga, saya baik. Bagaimana denganmu?",
      "timestamp": "2022-01-01T08:05:00.000Z",
      "status": "read",
      "attachment": null
    }
  ]
}
```

### group

```json
{
  "_id": "group_123",
  "messages": [
    {
      "_id": "message_1",
      "sender_id": "user_3",
      "receiver_ids": [
        {"user_id": "user_1", "status": "pending"},
        {"user_id": "user_2", "status": "pending"},
        {"user_id": "user_4", "status": "pending"}
      ],
      "content": "Bangun woy!!!",
      "timestamp": "2022-01-01T08:00:00.000Z",
      "status": "sent",
      "attachment": [{
        "file_name": "image1.jpg",
        "type": "image",
        "url": "https://dummyimage.com/600x400/000/fff&text=Bangun+woy!!!"
      }]
    },
    {
      "_id": "message_2",
      "sender_id": "user_2",
      "receiver_ids": [
        {"user_id": "user_1", "status": "read"},
        {"user_id": "user_3", "status": "read"},
        {"user_id": "user_4", "status": "read"}
      ],
      "content": "Selamat pagi! Sudah bangun, nih?",
      "timestamp": "2022-01-01T08:05:00.000Z",
      "status": "read",
      "attachment": null
    }
  ],
  "created_by": "user_3",
  "members": [
    {
      "user_id": "user_1",
      "role": "admin"
    },
    {
      "user_id": "user_2",
      "role": "member"
    },
    {
      "user_id": "user_3",
      "role": "admin"
    },
    {
      "user_id": "user_4",
      "role": "member"
    }
  ]
}
```

