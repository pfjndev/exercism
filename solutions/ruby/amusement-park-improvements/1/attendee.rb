class Attendee
  def initialize(height)
    @height = height
  end

  def issue_pass!(pass_id)
    @pass_id = pass_id
  end

  def revoke_pass!
    @pass_id = nil
  end

  # Do not edit above methods, add your own methods below.

  def has_pass?
    @pass_id != nil ? true : false
  end

  def fits_ride?(ride_minimum_height)
    return false unless @height >= ride_minimum_height
      true
  end

  def allowed_to_ride?(ride_minimum_height)
    return false unless fits_ride?(ride_minimum_height) && has_pass?
      true
  end
end
