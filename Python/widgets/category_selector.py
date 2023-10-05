# from document_types import xxx

def get_category():
    user_input = input("Press select the category of the document that you want to create (1-4)\n1.Experimental log\n")
    try:
        category = user_input
        print("You selected:", user_input, "\n==========================")
        return category
    except ValueError:
        print("Invalid selection. Try again.")    