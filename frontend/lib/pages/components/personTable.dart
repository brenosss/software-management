import 'package:flutter/material.dart';
import 'package:frontend/entities/person.dart';

class PersonTable extends StatelessWidget {
  final List<Person> personList;
  PersonTable({this.personList});

  @override
  Widget build(BuildContext context) {
    return DataTable(
      showCheckboxColumn: false,
      columns: const <DataColumn>[
        DataColumn(
          label: Text(
            'Name',
            style: TextStyle(fontStyle: FontStyle.italic),
          ),
        ),
        DataColumn(
          label: Text(
            'Email',
            style: TextStyle(fontStyle: FontStyle.italic),
          ),
        ),
        DataColumn(
          label: Text(
            'Birthday',
            style: TextStyle(fontStyle: FontStyle.italic),
          ),
        ),
      ],
      rows: personList
          .map((personItem) => DataRow(
                onSelectChanged: (newValue) {
                  print(personItem.personId);
                },
                cells: <DataCell>[
                  DataCell(Text(personItem.name)),
                  DataCell(Text(personItem.email)),
                  DataCell(Text(personItem.birthday)),
                ],
              ))
          .toList(),
    );
  }
}
