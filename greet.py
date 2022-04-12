def greeter(name: str, message: str) -> str:
    return f"Hi {name}! {message}"


if __name__ == "__main__":
    message = "You are welcome to Paradise on Earth!"
    greeting = greeter("Adam", message)
    print(greeting)