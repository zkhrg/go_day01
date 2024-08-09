import json
import xml.etree.ElementTree as ET

def json_to_xml(json_obj):
    def create_ingredients_xml(ingredients):
        ingredients_elem = ET.Element('ingredients')
        for ingredient in ingredients:
            item_elem = ET.Element('item')
            
            itemname_elem = ET.SubElement(item_elem, 'itemname')
            itemname_elem.text = ingredient.get('ingredient_name', '')
            
            itemcount_elem = ET.SubElement(item_elem, 'itemcount')
            itemcount_elem.text = str(ingredient.get('ingredient_count', ''))
            
            itemunit_elem = ET.SubElement(item_elem, 'itemunit')
            itemunit_elem.text = ingredient.get('ingredient_unit', '')
            
            ingredients_elem.append(item_elem)
        return ingredients_elem
    
    recipes_elem = ET.Element('recipes')
    
    for cake in json_obj['cake']:
        cake_elem = ET.Element('cake')
        
        name_elem = ET.SubElement(cake_elem, 'name')
        name_elem.text = cake.get('name', '')
        
        stovetime_elem = ET.SubElement(cake_elem, 'stovetime')
        stovetime_elem.text = cake.get('time', '')
        
        ingredients_elem = create_ingredients_xml(cake.get('ingredients', []))
        cake_elem.append(ingredients_elem)
        
        recipes_elem.append(cake_elem)
    
    return ET.tostring(recipes_elem, encoding='unicode', method='xml')

# Sample JSON data
json_data = '''

'''

# Convert JSON to Python dict
json_obj = json.loads(json_data)

# Convert to XML
xml_result = json_to_xml(json_obj)

# Print XML
print(xml_result)