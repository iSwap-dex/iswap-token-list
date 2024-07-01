### 链信息配置
`chains.json`

例子
```json
{
  "ethereum": 1,
  "sepolia": 11155111,
  "polygon": 137,
  "arbitrum": 42161,
  "arbitrum_nova": 42170,
  "optimism": 10,
  "optimism-sepolia": 11155420,
  "base": 8453,
  "base-sepolia": 84532,
  "bitlayer": 200901,
  "bitlayer-testnet": 200810,
  "cyber": 7560,
  "linea": 59144,
  "mode": 34443,
  "pgn": 424
}
```

### 增加 token 和 icon
##### 添加 icon 
`data/{token}/logo.svg` or `data/{token}/logo.png`
##### 添加 token 信息
`data/{token}/data.json`

例子
```json
{
    "name": "1INCH Token",
    "symbol": "1INCH",
    "decimals": 18,
    "website": "https://1inch.io/",
    "twitter": "@1inch",
    "tokens": {
      "ethereum": {
        "address": "0x111111111117dC0aa78b770fA6A738034120C302"
      },
      "optimism": {
        "address": "0xAd42D013ac31486B73b6b059e748172994736426"
      },
      "base": {
        "address": "0xc5fecC3a29Fb57B5024eEc8a2239d4621e111CBE"
      }
    }
}
```


### 获取 token.list.json 
<img width="1743" alt="image" src="https://github.com/iSwap-dex/iswap-token-list/assets/4955106/de5ce9c6-c969-454d-8ccb-db7cc5dffee4">


### 链信息配置 tokenlist

<img width="906" alt="image" src="https://github.com/iSwap-dex/iswap-token-list/assets/4955106/3eb75ac0-269a-49a7-81b9-2564a67c5048">



### 添加 token.list.json
<img width="1061" alt="image" src="https://github.com/iSwap-dex/iswap-token-list/assets/4955106/a16613a9-1859-4094-be41-37f7d2c9e66d">
