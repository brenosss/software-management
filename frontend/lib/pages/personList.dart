import 'package:flutter/material.dart';
import 'package:frontend/entities/person.dart';
import 'package:frontend/pages/components/personTable.dart';
import 'package:frontend/services/person.dart';
import 'package:frontend/pages/components/sideBar.dart';

class PersonList extends StatefulWidget {
  PersonList({Key key}) : super(key: key);

  @override
  _PersonListScreenState createState() => _PersonListScreenState();
}

class _PersonListScreenState extends State<PersonList> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("List Person"),
      ),
      drawer: SideBar(),
      body: Center(
        child: FutureBuilder<List<Person>>(
          future: fetchPersonList(),
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return PersonTable(personList: snapshot.data);
            } else if (snapshot.hasError) {
              return Text("${snapshot.error}");
            }
            // By default, show a loading spinner.
            return CircularProgressIndicator();
          },
        ),
      ),
    );
  }
}
