import { RSABSSA } from "@cloudflare/blindrsa-ts";
import { Base64 } from 'js-base64';


const variants = [
    RSABSSA.SHA384.PSS.Randomized,
    RSABSSA.SHA384.PSSZero.Randomized,
    RSABSSA.SHA384.PSS.Deterministic,
    RSABSSA.SHA384.PSSZero.Deterministic,
];

const suite = RSABSSA.SHA384.PSS.Randomized();

function convertPemToBinary(pem: string) {
    var lines = pem.split('\n')
    var encoded = ''
    for (var i = 0; i < lines.length; i++) {
        if (lines[i].trim().length > 0 &&
            lines[i].indexOf('-BEGIN RSA PRIVATE KEY-') < 0 &&
            lines[i].indexOf('-BEGIN RSA PUBLIC KEY-') < 0 &&
            lines[i].indexOf('-END RSA PRIVATE KEY-') < 0 &&
            lines[i].indexOf('-END RSA PUBLIC KEY-') < 0) {
            encoded += lines[i].trim()
        }
    }
    return Base64.toUint8Array(encoded)
}

function importPublicKey(pemKey: string) {
    return new Promise<CryptoKey>(function (resolve) {
        let importer = crypto.subtle.importKey("spki", convertPemToBinary(pemKey), {
            "name": "RSA-PSS",
            "hash": {
                "name": "SHA-384"
            }
        }, true, ["verify"])
        importer.then(function (key) {
            resolve(key)
        })
    })
}

async function getJSON(endpoint: string) {
    let myObject = await fetch(endpoint);
    let myText = await myObject.text();
    let myJson = await JSON.parse(myText);
    return myJson
}

async function getPublicKey() {

    let pk = await getJSON("/api/json/publickey");
    return pk
}

async function postJSON(endpoint: string, json: object) {

    let myRequest = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify(json)
    });
    return
}

async function main() {

    let pemPublicKey = await getPublicKey();
    let publicKey = await importPublicKey(pemPublicKey.serverpublickey)
    console.log(publicKey)

    const msgString = 'Alice and Bob';
    const message = new TextEncoder().encode(msgString);
    const preparedMsg = suite.prepare(message);
    const { blindedMsg, inv } = await suite.blind(publicKey, preparedMsg);

    console.log(blindedMsg)

    await postJSON("/api/json/reqBlindSignature", {"blindedMsg":Base64.fromUint8Array(blindedMsg)})

}
main();