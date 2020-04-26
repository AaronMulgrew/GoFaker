# test
import json

# with open('names.json') as json_file:
	# names = json_file.read()
	# names = names.replace("'","\"")
	# with open("output.json", "w") as output:
		# output.write(names)


with open('names.json') as json_file:
	with open('firstnames.go') as go_file:
		go_file = go_file.read()
		data = json.load(json_file)
		#print(data)
		index = go_file.find('/// put it here')
		first_half_of_go_file = go_file[:index]
		second_half_of_go_file = go_file[index+15:]
		allItems = ""
		for i in range(0, len(data)):
			item = data[i]
			allItems = allItems + '	allNames[' + str(i) + '] = "' + item.title().strip() + '"' + '\n'

	with open('firstnames.go', 'w') as output_file:
		output_file.write(first_half_of_go_file + allItems + second_half_of_go_file)