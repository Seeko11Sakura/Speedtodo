class Task {
  final int? id;
  final String title;
  final String notes;
  final String status;
  final int priority;
  final int? listId;
  final String? dueAt;
  final String? remindAt;
  final bool allDay;
  final String repeatRule;
  final String? completedAt;
  final String? createdAt;

  Task({
    this.id,
    this.title = '',
    this.notes = '',
    this.status = 'active',
    this.priority = 0,
    this.listId,
    this.dueAt,
    this.remindAt,
    this.allDay = false,
    this.repeatRule = '',
    this.completedAt,
    this.createdAt,
  });

  factory Task.fromJson(Map<String, dynamic> json) => Task(
        id: json['id'] as int?,
        title: json['title'] ?? '',
        notes: json['notes'] ?? '',
        status: json['status'] ?? 'active',
        priority: json['priority'] ?? 0,
        listId: json['listId'] as int?,
        dueAt: json['dueAt'] as String?,
        remindAt: json['remindAt'] as String?,
        allDay: json['allDay'] ?? false,
        repeatRule: json['repeatRule'] ?? '',
        completedAt: json['completedAt'] as String?,
        createdAt: json['createdAt'] as String?,
      );

  Map<String, dynamic> toJson() => {
        'id': id,
        'title': title,
        'notes': notes,
        'status': status,
        'priority': priority,
        'listId': listId,
        'dueAt': dueAt,
        'remindAt': remindAt,
        'allDay': allDay,
        'repeatRule': repeatRule,
        'completedAt': completedAt,
        'createdAt': createdAt,
      };
}
