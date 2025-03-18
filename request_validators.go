package unsend

type ValidationError struct {
	Errors []string
}

func (req CreateContactRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.ContactBookId == "" {
		errors.Errors = append(errors.Errors, "'ContactBookId' is required")
	}
	if req.Email == "" {
		errors.Errors = append(errors.Errors, "'Email' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req UpdateContactRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.ContactBookId == "" {
		errors.Errors = append(errors.Errors, "'ContactBookId' is required")
	}
	if req.ContactId == "" {
		errors.Errors = append(errors.Errors, "'ContactId' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req UpsertContactRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.ContactBookId == "" {
		errors.Errors = append(errors.Errors, "'ContactBookId' is required")
	}

	if req.ContactId == "" {
		errors.Errors = append(errors.Errors, "'ContactId' is required")
	}

	if req.Email == "" {
		errors.Errors = append(errors.Errors, "'Email' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req DeleteContactRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.ContactBookId == "" {
		errors.Errors = append(errors.Errors, "'ContactBookId' is required")
	}

	if req.ContactId == "" {
		errors.Errors = append(errors.Errors, "'ContactId' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req GetContactRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.ContactBookId == "" {
		errors.Errors = append(errors.Errors, "'ContactBookId' is required")
	}

	if req.ContactId == "" {
		errors.Errors = append(errors.Errors, "'ContactId' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req GetEmailRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.EmailId == "" {
		errors.Errors = append(errors.Errors, "'EmailId' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req SendEmailRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if len(req.To) == 0 {
		errors.Errors = append(errors.Errors, "'To' is required")
	}

	if len(req.From) == 0 {
		errors.Errors = append(errors.Errors, "'From' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req UpdateScheduleRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.EmailId == "" {
		errors.Errors = append(errors.Errors, "'EmailId' is required")
	}

	if req.ScheduledAt == "" {
		errors.Errors = append(errors.Errors, "'ScheduledAt' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

func (req CancelScheduleRequest) Validate() *ValidationError {
	errors := new(ValidationError)
	if req.EmailId == "" {
		errors.Errors = append(errors.Errors, "'EmailId' is required")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}
