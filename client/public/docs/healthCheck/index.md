# Health Check
ヘルスチェック機能を有効にしたサービスでは、<かっこいい名前>がサービスの死活管理を行うようになります。
例えば、9090ポートで待ち受けているサービスにが何らかの原因で停止している場合、Webhook経由で管理者に通知することができます。

## CD機能との連携
ヘルスチェック機能とCD機能の双方を有効にしているサービスでは、停止しているサービスについて再起動をする事ができます。
これにより、手動でサービスを再起動する必要がなくなります。
