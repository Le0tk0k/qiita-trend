# Qiita Trend
Qiitaのトレンド記事を取得するAPI

## Explanation Entry
[GoでQiitaのトレンドを取得するAPI①作ってみた](https://qiita.com/Le0tk0k/items/5ce324dba27c892d9f28)

## Install

```bash
$ go get github.com/Le0tk0k/qiita-trend
```

## Usage

```bash
$ go run main.go
$ curl http://localhost:8080/trend
```

## Example
返ってくるJSONの例です。
※見やすくするため整形しています。

```bash
{
    "trend": {
        "edges": [
            {
                "followingLikers": [],
                "isLikedByViewer": false,
                "isNewArrival": false,
                "hasCodeBlock": true,
                "node": {
                    "createdAt": "2020-09-11T07:53:22Z",
                    "likesCount": 338,
                    "title": "オブジェクト指向は単なる【整理術】だよ",
                    "uuid": "2aa9859d4da41991bbb7",
                    "author": {
                        "profileImageUrl": "https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/163629/profile-images/1567829792",
                        "urlName": "br_branch"
                    }
                }
            },
            
            省略
             .            
             .
             .
             .
            省略

            {
                "followingLikers": [],
                "isLikedByViewer": false,
                "isNewArrival": true,
                "hasCodeBlock": false,
                "node": {
                    "createdAt": "2020-09-13T09:55:28Z",
                    "likesCount": 15,
                    "title": "micro:bitプログラミング入門講座テキスト",
                    "uuid": "df4b502b43fde7eba35b",
                    "author": {
                        "profileImageUrl": "https://qiita-image-store.s3.amazonaws.com/0/26536/profile-images/1473684616",
                        "urlName": "noanoa07"
                    }
                }
            }
        ]
    },
    "scope": "daily"
}
```
