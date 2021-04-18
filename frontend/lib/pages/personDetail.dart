import 'package:flutter/material.dart';
import 'package:frontend/pages/components/sideBar.dart';
import 'package:frontend/services/person.dart';
import 'package:frontend/entities/person.dart';
import 'package:frontend/entities/personSkill.dart';

class PersonDetailArguments {
  final String personId;

  PersonDetailArguments(this.personId);
}

class PersonDetail extends StatefulWidget {
  PersonDetail({Key key}) : super(key: key);

  @override
  _PersonDetailScreenState createState() => _PersonDetailScreenState();
}

class _PersonDetailScreenState extends State<PersonDetail> {
  @override
  Widget build(BuildContext context) {
    final PersonDetailArguments args =
        ModalRoute.of(context).settings.arguments as PersonDetailArguments;
    return Scaffold(
      appBar: AppBar(
        title: Text("List Person"),
      ),
      drawer: SideBar(),
      body: Container(
        child: FutureBuilder<Person>(
          future: fetchPersonDetail(args.personId),
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return PersonInfo(snapshot.data);
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

class PersonInfo extends StatelessWidget {
  final Person person;
  PersonInfo(this.person);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(40.0),
      child: Container(
        decoration: BoxDecoration(
          color: Colors.cyan.shade50,
          borderRadius: BorderRadius.circular(10),
          boxShadow: [
            BoxShadow(
              color: Colors.grey,
              blurRadius: 2,
              offset: Offset(4, 8), // Shadow position
            ),
          ],
        ),
        child: Column(children: [
          Flexible(
            flex: 1,
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Row(
                children: <Widget>[
                  Expanded(
                      flex: 2,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          CircleAvatar(
                              radius: 60.0,
                              child: ClipOval(
                                child: Image.network(
                                  "https://dt2sdf0db8zob.cloudfront.net/wp-content/uploads/2019/12/9-Best-Online-Avatars-and-How-to-Make-Your-Own-for-Free-image1-5.png",
                                  fit: BoxFit.cover,
                                ),
                              )),
                        ],
                      )),
                  Expanded(
                      flex: 6,
                      child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: <Widget>[
                            Text(person.name,
                                style: TextStyle(
                                    color: Colors.black, fontSize: 40)),
                            Text("Phone: ${person.phone}",
                                style: TextStyle(
                                    color: Colors.black54, fontSize: 25)),
                            Text("Email: ${person.email}",
                                style: TextStyle(
                                    color: Colors.black54, fontSize: 25)),
                          ])),
                  Expanded(
                      flex: 2,
                      child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            ElevatedButton.icon(
                              onPressed: () {},
                              label: Text('Hire'),
                              icon: Icon(Icons.plus_one),
                            )
                          ]))
                ],
              ),
            ),
          ),
          Flexible(
              flex: 2,
              child: GridView.count(
                  padding: const EdgeInsets.all(50),
                  crossAxisCount: 4,
                  childAspectRatio: 2,
                  children: person.personSkills
                      .map((personSkill) =>
                          PersonSkillItem(personSkill: personSkill))
                      .toList()))
        ]),
      ),
    );
  }
}

class PersonSkillItem extends StatelessWidget {
  final PersonSkill personSkill;
  PersonSkillItem({this.personSkill});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Container(
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(10),
          boxShadow: [
            BoxShadow(
              color: Colors.grey,
              blurRadius: 2,
              offset: Offset(4, 8), // Shadow position
            ),
          ],
        ),
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            children: <Widget>[
              ListTile(
                leading: Icon(Icons.code),
                title: Text(personSkill.skill.name),
              ),
              Text("Type: ${personSkill.skill.type}",
                  style: TextStyle(color: Colors.grey)),
              Text("Value: ${personSkill.value}",
                  style: TextStyle(color: Colors.grey)),
              Text("Progress: ${personSkill.progress}",
                  style: TextStyle(color: Colors.grey)),
            ],
          ),
        ),
      ),
    );
  }
}
