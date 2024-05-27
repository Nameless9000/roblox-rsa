n = int("5f85a133692139651055c4d4cf04d7d66a3c002a3779ec6746a6b9fa558514653de0cd333b1796a8a5eb9e3eefc92a402af1f080447ae498a93a1923a7bd8853", 16)
e = int("10001", 16)

# example
headers = {"CM2-User":"5154741620","CM2-Nonce":"145328524","CM2-Signature":"16a576cca2a89ad303a0900c61502da4986c39fe37e692f7df1c3580c18b041f3cbaf88d1eb9637df1ead4266fbdcb4b416321fbae59a1773b5394ce8e8e9a87"}

signature = int(headers["CM2-Signature"], 16)

expected_userid = pow(signature, e, n) - int(headers["CM2-Nonce"])
user_id = int(headers["CM2-User"])

print(expected_userid == user_id)