import 'package:flutter/material.dart';

class InboxPage extends StatelessWidget {
  const InboxPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Inbox')),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {},
        label: const Text('Add Task'),
        icon: const Icon(Icons.add),
      ),
      body: const Center(child: Text('Your inbox is empty')),
    );
  }
}
