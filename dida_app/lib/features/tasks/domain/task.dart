class Task {
  const Task({
    required this.id,
    required this.title,
    required this.status,
    this.dueAt,
  });

  final int id;
  final String title;
  final String status;
  final DateTime? dueAt;
}
