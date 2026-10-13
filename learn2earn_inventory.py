

resources = [
    {"id": "R001", "name": "Laptop", "category": "Electronics", "total": 10, "available": 10},
    {"id": "R002", "name": "Keyboard", "category": "Accessories", "total": 5, "available": 5},
    {"id": "R003", "name": "Headset", "category": "Accessories", "total": 3, "available": 3}
]
fellows = {"F001": "Ada", "F002": "John", "F003": "Grace"}
borrow_records = []

LOW_STOCK_LIMIT = 4



def clean_id(text):
    return text.strip().upper()


def find_resource(resource_id):
    for item in resources:
        if item["id"] == resource_id:
            return item
    return None


def parse_positive_int(text):
    try:
        value = int(text.strip())
    except ValueError:
        return None
    if value <= 0:
        return None
    return value


def print_resources(items):
    if not items:
        print("No resources to show.")
        return
    print(f"{'ID':<8}{'Name':<15}{'Category':<15}{'Total':>6}{'Available':>11}")
    print("-" * 55)
    for r in items:
        print(f"{r['id']:<8}{r['name']:<15}{r['category']:<15}{r['total']:>6}{r['available']:>11}")


def units_on_loan(resource_id):
    return sum(rec["amount"] - rec["returned"]
               for rec in borrow_records if rec["resource_id"] == resource_id)



def add_resource():
    new_id = clean_id(input("New resource ID: "))
    if new_id == "":
        print("Error: ID cannot be blank.")
        return
    if find_resource(new_id) is not None:
        print(f"Error: resource ID {new_id} already exists.")
        return

    name = input("Name: ").strip()
    if name == "":
        print("Error: name cannot be blank.")
        return

    category = input("Category: ").strip()
    if category == "":
        print("Error: category cannot be blank.")
        return

    total = parse_positive_int(input("Total units: "))
    if total is None:
        print("Error: total units must be a positive whole number.")
        return

    resources.append({"id": new_id, "name": name, "category": category,
                      "total": total, "available": total})
    print(f"Added {name} ({new_id}) with {total} units.")


def list_resources():
    print_resources(resources)



def borrow_resource():
    fellow_id = clean_id(input("Fellow ID: "))
    if fellow_id not in fellows:
        print("Error: fellow ID not found.")
        return

    resource_id = clean_id(input("Resource ID: "))
    item = find_resource(resource_id)
    if item is None:
        print("Error: resource ID not found.")
        return

    quantity = parse_positive_int(input("Quantity: "))
    if quantity is None:
        print("Error: quantity must be a positive whole number.")
        return

    if quantity > item["available"]:
        print(f"Error: only {item['available']} {item['name']}(s) available, "
              f"you asked for {quantity}.")
        return

    item["available"] -= quantity
    borrow_records.append({
        "fellow_id": fellow_id,
        "fellow_name": fellows[fellow_id],
        "resource_id": item["id"],
        "resource_name": item["name"],
        "amount": quantity,
        "returned": 0
    })
    print(f"Success: {fellows[fellow_id]} borrowed {quantity} {item['name']}(s). "
          f"Available now: {item['available']}")



def return_resource():
    fellow_id = clean_id(input("Fellow ID: "))
    if fellow_id not in fellows:
        print("Error: fellow ID not found.")
        return

    resource_id = clean_id(input("Resource ID: "))
    item = find_resource(resource_id)
    if item is None:
        print("Error: resource ID not found.")
        return

    outstanding = sum(rec["amount"] - rec["returned"] for rec in borrow_records
                      if rec["fellow_id"] == fellow_id and rec["resource_id"] == resource_id)
    if outstanding == 0:
        print(f"Error: {fellows[fellow_id]} has no {item['name']}(s) on loan.")
        return

    quantity = parse_positive_int(input("Quantity to return: "))
    if quantity is None:
        print("Error: quantity must be a positive whole number.")
        return

    if quantity > outstanding:
        print(f"Error: {fellows[fellow_id]} only has {outstanding} {item['name']}(s) "
              f"on loan, cannot return {quantity}.")
        return

    remaining = quantity
    for rec in borrow_records:
        if rec["fellow_id"] == fellow_id and rec["resource_id"] == resource_id:
            take = min(remaining, rec["amount"] - rec["returned"])
            rec["returned"] += take
            remaining -= take
            if remaining == 0:
                break

    item["available"] += quantity
    print(f"Success: {fellows[fellow_id]} returned {quantity} {item['name']}(s). "
          f"Available now: {item['available']}")



def search_by_name():
    term = input("Search name: ").strip().lower()
    if term == "":
        print("Error: search text cannot be blank.")
        return
    matches = [r for r in resources if term in r["name"].lower()]
    if not matches:
        print(f"No resources match '{term}'.")
        return
    print_resources(matches)


def filter_by_category():
    categories = sorted({r["category"] for r in resources})
    print("Categories:", ", ".join(categories))
    term = input("Category: ").strip().lower()
    matches = [r for r in resources if r["category"].lower() == term]
    if not matches:
        print(f"No resources in category '{term}'.")
        return
    print_resources(matches)



def generate_report():
    total_units = sum(r["total"] for r in resources)
    available_units = sum(r["available"] for r in resources)
    borrowed_units = sum(units_on_loan(r["id"]) for r in resources)

    print("===== INVENTORY REPORT =====")
    print(f"Total units:      {total_units}")
    print(f"Available units:  {available_units}")
    print(f"Borrowed units:   {borrowed_units}")
    if total_units != available_units + borrowed_units:
        print("WARNING: total does not equal available + borrowed!")

    print(f"\nLow stock (fewer than {LOW_STOCK_LIMIT} available):")
    low = [r for r in resources if r["available"] < LOW_STOCK_LIMIT]
    if low:
        for r in low:
            print(f"  {r['name']} ({r['available']})")
    else:
        print("  None")

    print("\nMost borrowed resource:")
    counts = {r["id"]: units_on_loan(r["id"]) for r in resources}
    highest = max(counts.values()) if counts else 0
    if highest == 0:
        print("  No resources are currently borrowed.")
    else:
        leaders = [r for r in resources if counts[r["id"]] == highest]
        for r in leaders:
            print(f"  {r['name']} ({highest})")
        if len(leaders) > 1:
            print("  (tie)")



def main():
    while True:
        print("\n===== LEARN2EARN INVENTORY =====")
        print("1. List resources")
        print("2. Add resource")
        print("3. Borrow resource")
        print("4. Return resource")
        print("5. Search by name")
        print("6. Filter by category")
        print("7. Generate report")
        print("0. Exit")
        choice = input("Choose an option: ").strip()

        if choice == "1":
            list_resources()
        elif choice == "2":
            add_resource()
        elif choice == "3":
            borrow_resource()
        elif choice == "4":
            return_resource()
        elif choice == "5":
            search_by_name()
        elif choice == "6":
            filter_by_category()
        elif choice == "7":
            generate_report()
        elif choice == "0":
            print("Goodbye.")
            break
        else:
            print("Invalid choice. Please enter a number from 0 to 7.")


if __name__ == "__main__":
    main()
