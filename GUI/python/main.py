import tkinter as tk
from tkinter import filedialog, messagebox, ttk
import datetime

class MarkdownEditor:
    def __init__(self, root):
        self.root = root
        root.title("Markdown File Generator with YAML Front Matter")
        root.configure(bg='#f0f0f0')

        # Styling
        style = ttk.Style()
        style.configure('TLabel', background='#f0f0f0', font=('Arial', 10))
        style.configure('TButton', font=('Arial', 10))
        style.configure('TEntry', font=('Arial', 10))

        # Date Field
        frame_date = ttk.Frame(root)
        frame_date.pack(padx=10, pady=5, fill='x')

        ttk.Label(frame_date, text="Date:").pack(side='left', padx=(0, 10))
        self.date_entry = ttk.Entry(frame_date)
        self.date_entry.insert(0, datetime.datetime.now().strftime("%Y-%m-%d"))
        self.date_entry.pack(side='right', expand=True, fill='x')

        # Category Dropdown
        frame_category = ttk.Frame(root)
        frame_category.pack(padx=10, pady=5, fill='x')

        ttk.Label(frame_category, text="Category:").pack(side='left', padx=(0, 10))
        self.category_var = tk.StringVar(root)
        self.category_var.set("Experimental log")  # default value
        self.category_options = ["Experimental log", "Discussion", "Reading", "Inspiration"]
        self.category_menu = ttk.OptionMenu(frame_category, self.category_var, *self.category_options)
        self.category_menu.pack(side='right', expand=True, fill='x')

        # Author Field
        frame_author = ttk.Frame(root)
        frame_author.pack(padx=10, pady=5, fill='x')

        ttk.Label(frame_author, text="Author:").pack(side='left', padx=(0, 10))
        self.author_entry = ttk.Entry(frame_author)
        self.author_entry.insert(0, "Jie Hua")
        self.author_entry.pack(side='right', expand=True, fill='x')

        # Email Field
        frame_email = ttk.Frame(root)
        frame_email.pack(padx=10, pady=5, fill='x')

        ttk.Label(frame_email, text="Email:").pack(side='left', padx=(0, 10))
        self.email_entry = ttk.Entry(frame_email)
        self.email_entry.insert(0, "Jie.Hua@lmu.de")
        self.email_entry.pack(side='right', expand=True, fill='x')

        # Text Area for Markdown Content
        self.text_area = tk.Text(root, height=20, width=60, font=('Arial', 10), padx=10, pady=10)
        self.text_area.pack(padx=10, pady=5, expand=True, fill='both')

        # Save Button
        self.save_button = ttk.Button(root, text="Save as Markdown", command=self.save_as_markdown)
        self.save_button.pack(padx=10, pady=(0, 10))

    def save_as_markdown(self):
        # Collecting all the inputs
        date = self.date_entry.get()
        category = self.category_var.get()
        author = self.author_entry.get()
        email = self.email_entry.get()
        text_content = self.text_area.get("1.0", "end-1c")

        # Formatting the YAML front matter
        yaml_content = f"---\nDate: {date}\nCategory: {category}\nAuthor: {author}\nEmail: {email}\n---\n\n"

        # Complete Markdown content with YAML front matter
        complete_content = yaml_content + text_content

        # Constructing default file name
        default_filename = f"{category}_{date}.md"
        file_path = filedialog.asksaveasfilename(initialfile=default_filename, defaultextension=".md", 
                                             filetypes=[("Markdown files", "*.md")])
        if file_path:
            with open(file_path, "w") as file:
                file.write(complete_content)
            messagebox.showinfo("Success", "File saved successfully.")

if __name__ == "__main__":
    root = tk.Tk()
    app = MarkdownEditor(root)
    root.geometry("600x500")  # Set initial size of the window
    root.mainloop()
