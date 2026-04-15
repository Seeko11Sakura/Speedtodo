package com.speedtodo.service;
import com.speedtodo.model.FocusSession;
import com.speedtodo.repository.FocusSessionRepository;
import org.springframework.stereotype.Service;
import java.time.LocalDateTime;
import java.util.List;
@Service
public class FocusSessionService {
    private final FocusSessionRepository repository;
    public FocusSessionService(FocusSessionRepository repository) { this.repository = repository; }
    public List<FocusSession> list() { return repository.findAll(); }
    public FocusSession create(FocusSession s) { s.setStatus("running"); s.setStartedAt(LocalDateTime.now()); return repository.save(s); }
    public FocusSession complete(Long id) { FocusSession s = repository.findById(id).orElseThrow(() -> new RuntimeException("Not found")); s.setStatus("completed"); s.setEndedAt(LocalDateTime.now()); return repository.save(s); }
    public FocusSession cancel(Long id) { FocusSession s = repository.findById(id).orElseThrow(() -> new RuntimeException("Not found")); s.setStatus("cancelled"); s.setEndedAt(LocalDateTime.now()); return repository.save(s); }
}