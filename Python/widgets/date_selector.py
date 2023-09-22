import datetime

# Print date of today and let users to select the date they will use.

def get_date():
    date_format = "%Y-%m-%d"
    date_system = datetime.datetime.now().strftime(date_format)
    print("Guten Tag! Today is "+date_system+"\n==========================")
    while True:
        user_input = input("Press Enter to use the current system date, or enter a date (format: YYYY-MM-DD): ")
        if not user_input:  # User presses Enter
            current_date = datetime.datetime.now()
            print("The current system date is:", current_date.strftime("%Y-%m-%d"),"\n==========================")
            return current_date
        else:
            try:
                user_date = datetime.datetime.strptime(user_input, "%Y-%m-%d")
                print("You entered the date:", user_date.strftime("%Y-%m-%d"),"\n==========================")
                return user_date
            except ValueError:
                print("Invalid date format. Please use the format YYYY-MM-DD. Try again.")         
