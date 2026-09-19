import json
import sys
import urllib.error
import urllib.request


def post_json(url, payload):
    data = json.dumps(payload).encode("utf-8")

    request = urllib.request.Request(
        url,
        data=data,
        headers={
            "Content-Type": "application/json",
        },
        method="POST",
    )

    try:
        with urllib.request.urlopen(request) as response:
            body = response.read().decode("utf-8")

            print(f"Status: {response.status}")
            print(body)

    except urllib.error.HTTPError as error:
        print(f"HTTP Error: {error.code}")
        print(error.read().decode("utf-8"))

    except urllib.error.URLError as error:
        print(f"Connection Error: {error.reason}")


def main():
    base_url = "http://localhost"

    if len(sys.argv) > 1:
        base_url = sys.argv[1].rstrip("/")

    payload = {
        "tenant": "demoA",
        "source": "api",
        "event_type": "sample_api_event",
        "user": "sample-user",
        "ip": "203.0.113.10",
        "reason": "sample_script",
    }

    print(f"Sending log to {base_url}/api/ingest")

    post_json(
        f"{base_url}/api/ingest",
        payload,
    )


if __name__ == "__main__":
    main()