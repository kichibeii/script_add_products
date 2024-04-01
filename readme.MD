### Product Syn Script
- mock_database : for mocking creating table mysql
- server : main script to sync product from database to shopify


### Step to sync
1. Get All data from database that have empty shopify_id 
    - create some mapping id and title for update later
2. Preparing StaggedInput (shopify)
3. Upload File to shopify (still got an error)
4. BulkdOperationInput 
5. waiting for BulkdOperationInput to success
6. if success get the url and download the title
7. Checking map for checking title that created
8. Update table products that already create for handling not duplicate input shopify
