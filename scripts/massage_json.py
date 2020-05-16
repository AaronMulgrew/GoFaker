# test
import json

# with open('names.json') as json_file:
	# names = json_file.read()
	# names = names.replace("'","\"")
	# with open("output.json", "w") as output:
		# output.write(names)


with open('media_types.json') as json_file:
	with open('server.go') as go_file:
		go_file = go_file.read()
		data = json.load(json_file)
		#print(data)
		index = go_file.find('/// put it here')
		first_half_of_go_file = go_file[:index]
		second_half_of_go_file = go_file[index+15:]
		allItems = ""
		i = 0
		for extension, content_type in data.items():
			#print(key)
			#print(value)
			allItems = allItems + '	allFileTypes[' + str(i) + '] = ' + "FileType{\"" + extension + '","' + content_type + '"}' + '\n'
			i = i + 1
		print(allItems)
	with open('server.go', 'w') as output_file:
		output_file.write(first_half_of_go_file + allItems + second_half_of_go_file)