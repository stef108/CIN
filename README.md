sequenceDiagram
    participant Dev as Developer
    participant GH as GitHub Repo
    participant GA as GitHub Actions
    participant Docker as Docker Engine

    Note over Dev, Docker: Requirement 1.2: Auto-trigger on push

    Dev->>GH: Push Code (git push)
    GH->>GA: Trigger Pipeline (Event: push)
    
    rect rgb(240, 248, 255)
        Note right of GA: Job 1: Quality Control
        GA->>GA: Checkout Code
        GA->>GA: Run Linter (Req 2.2)
        GA->>GA: Run Unit Tests (Req 2.1)
        
        alt Tests Fail
            GA-->>Dev: Email Notification (Req 2.1)
        else Tests Pass
            GA->>Docker: Start Build Job
        end
    end

    rect rgb(255, 240, 245)
        Note right of GA: Job 2: Packaging (Req 2.4)
        Docker->>Docker: Build Multi-stage Image
        Docker-->>GA: Image Created (devops-app:sha)
    end
