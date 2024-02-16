import os

def generate_files_in_folders(root_directory):
    for foldername, subfolders, _ in os.walk(root_directory):
        if ".git" in foldername:
            continue
        if foldername == "./template" or foldername == "./":
            continue        
        
        folder_path = os.path.join(root_directory, foldername)

        # Check if a file named after the folder exists
        folder_name = os.path.basename(folder_path)
        new_file_name = f"{folder_name.lower()}_test.go"
        new_file_path = os.path.join(folder_path, new_file_name.lower())
        
        if not os.path.exists(new_file_path):
            # Create the file with folder name appended if it doesn't exist
            with open(new_file_path, 'w') as new_file:
                new_file.write(f"package main")

# Replace 'root_directory' with the path to the root folder where you want to generate files
generate_files_in_folders('./')
